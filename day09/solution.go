// Advent of Code: Day 09
//
// https://adventofcode.com/2024/day/9
//
// Part 1 Idea: The first challenge is choosing a representation for the disk map.
// The naive approach is to represent the entire map in a slice, where each cell
// is the ID of the block in that position. A more compact idea is to represent
// the map as a slice, where each cell is a struct with the ID and how many memory
// blocks are taken up at that position. An ID of -1 could represent free memory.
// I believe this is the most space efficient we can be, since we need to at minimum
// know for each ID the blocks it takes up and where. Let's try that.
//
// Part 2 Idea: The representation I chose in part 1 makes this much easier. I'm going
// to try the naive approach first, which will start with the leftmost free block
// every time. Given that we need to move files to the leftmost memory block, not the
// smallest memory block that could fit it, the only way I can think to optimize is to
// keep a second list of only free blocks. This has effectively the same big-O complexity,
// and may not be worth the extra complexity.
// For now, we can fit it into our part 1 solution with very little extra code.
package day09

import (
	"advent/util"
	"bufio"
	"slices"
)

// A disk span repressents a contiguous span of memory blocks, where an ID is stored.
// An ID of -1 represents free memory.
type DiskSpan struct {
	Id    int
	Start int
	Size  int
}

type Day09Solution struct {
	diskMap []DiskSpan
}

func NewDay09Solution(filepath string) (*Day09Solution, error) {
	diskMap := getdiskMap(filepath)
	return &Day09Solution{diskMap}, nil
}

func (s *Day09Solution) PartOneAnswer() (int, error) {
	reindexedDiskMap := s.reindexFiles(s.diskMap, true)
	return s.getDiskMapChecksum(reindexedDiskMap), nil
}

func (s *Day09Solution) PartTwoAnswer() (int, error) {
	reindexedDiskMap := s.reindexFiles(s.diskMap, false)
	return s.getDiskMapChecksum(reindexedDiskMap), nil
}

// reindexFilesFragmented moves blocks around to compact the disk map, starting with
// files at the end of the disk. The process is repeated until blocks from
// left to right are used until there are no blocks left. All free blocks
// follow. If fragment is true, files can be split up. Otherwise they cannot.
func (s *Day09Solution) reindexFiles(diskMap []DiskSpan, fragment bool) []DiskSpan {
	reindexedDiskMap := make([]DiskSpan, len(diskMap))
	copy(reindexedDiskMap, s.diskMap)
	freeSpanIndex := 0
	usedSpanIndex := len(s.diskMap) - 1
	for usedSpanIndex > 0 && freeSpanIndex <= len(reindexedDiskMap) {
		if freeSpanIndex >= usedSpanIndex {
			if fragment {
				// there's no more free space to the left of used space
				return reindexedDiskMap
			} else {
				// we may just not have space for this block -- let's move on to the next
				usedSpanIndex--
				freeSpanIndex = 0
			}
		}
		if reindexedDiskMap[freeSpanIndex].Id != -1 {
			// not a free block, skip
			freeSpanIndex++
		} else if reindexedDiskMap[usedSpanIndex].Id == -1 {
			// not a used block, skip
			usedSpanIndex--
		} else {
			freeDiskSpan := reindexedDiskMap[freeSpanIndex]
			usedDiskSpan := reindexedDiskMap[usedSpanIndex]
			if reindexedDiskMap[freeSpanIndex].Size < reindexedDiskMap[usedSpanIndex].Size {
				// we can only fit part of the used block in the free block
				if fragment {
					reindexedDiskMap[freeSpanIndex].Id = usedDiskSpan.Id
					reindexedDiskMap[usedSpanIndex].Size -= freeDiskSpan.Size
				}
				freeSpanIndex++
			} else {
				// we can fit the entire used block in the free block
				leftover := DiskSpan{-1, freeDiskSpan.Start + usedDiskSpan.Size, freeDiskSpan.Size - usedDiskSpan.Size}
				// updating the free disk span to hold the used disk span and reflect the new size
				reindexedDiskMap[freeSpanIndex].Id = usedDiskSpan.Id
				reindexedDiskMap[freeSpanIndex].Size = usedDiskSpan.Size
				// updating the used disk span to show it is free
				reindexedDiskMap[usedSpanIndex].Id = -1
				// finally, inserting the leftover if necessary
				if leftover.Size > 0 {
					reindexedDiskMap = slices.Insert(reindexedDiskMap, freeSpanIndex+1, leftover)
				}
				if fragment {
					freeSpanIndex++
				} else {
					freeSpanIndex = 0
				}
				usedSpanIndex--
			}
		}
	}
	return reindexedDiskMap
}

func (s *Day09Solution) getDiskMapChecksum(diskMap []DiskSpan) int {
	checksum := 0
	for _, diskSpan := range diskMap {
		if diskSpan.Id != -1 {
			checksum += s.getDiskSpanChecksum(diskSpan)
		}
	}
	return checksum
}

func (s *Day09Solution) getDiskSpanChecksum(diskSpan DiskSpan) int {
	checksum := 0
	for i := diskSpan.Start; i < diskSpan.Start+diskSpan.Size; i++ {
		checksum += i * diskSpan.Id
	}
	return checksum
}

func getdiskMap(filepath string) []DiskSpan {
	diskMap := make([]DiskSpan, 0)
	util.ProcessFile(filepath, func(scanner *bufio.Scanner) error {
		isFile := true
		id := 0
		pos := 0
		for scanner.Scan() {
			line := scanner.Text()
			for _, r := range line {
				size := int(r - '0')
				if isFile {
					diskMap = append(diskMap, DiskSpan{id, pos, size})
					id++
				} else {
					diskMap = append(diskMap, DiskSpan{-1, pos, size})
				}
				isFile = !isFile
				pos += size
			}
		}
		return nil
	})
	return diskMap
}
