package blockset

var Lobby = [...]byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 33, 33, 33, 33, 1,
	1, 1, 1, 1, 1, 1, 1, 12, 13, 1, 1, 28, 29, 33, 33, 1,
	1, 1, 1, 1, 1, 1, 1, 10, 11, 1, 1, 26, 27, 33, 33, 36,
	37, 14, 15, 52, 53, 30, 31, 48, 21, 21, 21, 32, 32, 32, 32, 39,
	39, 36, 37, 55, 55, 52, 53, 21, 21, 21, 21, 32, 32, 32, 32, 14,
	15, 32, 32, 30, 31, 32, 32, 21, 49, 32, 32, 32, 32, 32, 32, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 3, 18, 19, 18, 19, 32,
	32, 32, 32, 32, 32, 32, 32, 4, 4, 4, 4, 20, 20, 20, 20, 32,
	32, 7, 8, 32, 32, 23, 24, 32, 32, 7, 8, 32, 32, 23, 24, 42,
	43, 44, 45, 58, 59, 60, 61, 64, 65, 66, 67, 80, 81, 82, 83, 32,
	32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 40, 40, 33, 33, 56, 56, 1,
	1, 1, 1, 1, 1, 1, 1, 6, 1, 1, 1, 22, 33, 33, 33, 1,
	1, 1, 1, 1, 1, 1, 1, 14, 15, 14, 15, 30, 31, 30, 31, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 38,
	39, 39, 39, 54, 57, 21, 21, 54, 57, 32, 32, 54, 57, 32, 32, 39,
	39, 39, 39, 21, 21, 21, 21, 32, 32, 32, 32, 32, 32, 32, 32, 34,
	35, 34, 35, 50, 51, 50, 51, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 54,
	57, 32, 32, 54, 57, 32, 32, 54, 57, 32, 32, 54, 57, 32, 32, 38,
	39, 39, 39, 48, 21, 21, 21, 32, 32, 32, 32, 32, 32, 32, 32, 39,
	39, 39, 41, 21, 21, 21, 49, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 54, 57, 32, 32, 54, 57, 32, 32, 54, 57, 32, 32, 54, 57, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 38, 41, 33, 33, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 33, 33, 38, 41, 32,
	32, 32, 32, 32, 32, 32, 32, 42, 43, 44, 45, 58, 59, 60, 61, 38,
	41, 38, 41, 34, 35, 34, 35, 34, 35, 34, 35, 50, 51, 50, 51, 68,
	68, 68, 68, 84, 84, 84, 84, 55, 69, 55, 69, 69, 55, 69, 55, 9,
	39, 39, 25, 54, 55, 55, 57, 70, 55, 55, 71, 85, 86, 87, 55, 64,
	65, 66, 67, 80, 81, 82, 83, 32, 32, 32, 32, 32, 32, 32, 32, 7,
	8, 55, 69, 23, 24, 69, 55, 7, 8, 55, 69, 23, 24, 69, 55, 55,
	69, 55, 69, 69, 55, 69, 55, 55, 69, 55, 69, 69, 55, 69, 55, 55,
	69, 7, 8, 69, 55, 23, 24, 55, 69, 7, 8, 69, 55, 23, 24, 1,
	1, 1, 1, 1, 1, 1, 1, 38, 41, 38, 41, 34, 35, 34, 35, 34,
	35, 34, 35, 50, 51, 50, 51, 55, 69, 55, 69, 69, 55, 69, 55, 1,
	1, 1, 1, 1, 1, 1, 1, 72, 73, 1, 1, 88, 89, 33, 33, 1,
	1, 1, 1, 1, 1, 1, 1, 46, 47, 1, 1, 33, 33, 33, 33, 39,
	39, 46, 47, 21, 21, 21, 49, 32, 32, 32, 32, 32, 32, 32, 32, 39,
	39, 39, 41, 21, 21, 54, 57, 32, 32, 54, 57, 32, 32, 54, 57, 55,
	69, 55, 69, 69, 55, 69, 55, 55, 69, 4, 4, 69, 55, 20, 20, 55,
	69, 55, 69, 69, 55, 69, 55, 4, 4, 55, 69, 20, 20, 69, 55, 1,
	1, 1, 1, 33, 33, 33, 33, 55, 69, 55, 69, 69, 55, 69, 55, 1,
	1, 62, 1, 33, 33, 63, 33, 55, 69, 55, 69, 69, 55, 69, 55, 1,
	1, 1, 1, 33, 33, 33, 33, 7, 8, 7, 8, 23, 24, 23, 24, 7,
	8, 7, 8, 23, 24, 23, 24, 9, 39, 39, 25, 54, 55, 55, 57, 1,
	1, 1, 1, 33, 33, 33, 33, 55, 69, 7, 8, 69, 55, 23, 24, 55,
	69, 55, 69, 69, 55, 69, 55, 39, 39, 39, 39, 21, 21, 21, 21, 1,
	1, 1, 1, 33, 33, 33, 33, 32, 32, 32, 32, 32, 32, 32, 32, 70,
	55, 55, 71, 85, 86, 87, 55, 7, 8, 7, 8, 23, 24, 23, 24, 1,
	1, 1, 1, 38, 41, 33, 33, 54, 57, 32, 32, 54, 57, 32, 32, 55,
	69, 55, 69, 69, 55, 69, 55, 39, 39, 39, 41, 21, 21, 21, 21, 1,
	1, 72, 73, 33, 33, 88, 89, 55, 69, 55, 69, 69, 55, 69, 55, 40,
	40, 1, 1, 56, 56, 33, 33, 55, 69, 55, 69, 69, 55, 69, 55, 79,
	79, 79, 76, 79, 79, 79, 76, 79, 79, 79, 76, 79, 79, 79, 76, 55,
	69, 55, 69, 69, 55, 69, 55, 68, 68, 68, 68, 84, 84, 84, 84, 75,
	76, 75, 76, 75, 76, 75, 76, 75, 76, 75, 76, 77, 78, 77, 78, 54,
	57, 54, 57, 54, 57, 54, 57, 54, 57, 54, 57, 54, 57, 54, 57, 55,
	69, 55, 69, 69, 55, 69, 55, 38, 41, 38, 41, 54, 57, 54, 57, 39,
	41, 55, 55, 54, 57, 55, 55, 38, 41, 55, 55, 34, 35, 91, 91, 34,
	35, 46, 47, 50, 51, 33, 33, 55, 69, 55, 69, 69, 55, 69, 55, 55,
	69, 55, 69, 69, 55, 69, 55, 4, 4, 4, 4, 20, 20, 20, 20, 39,
	39, 39, 39, 21, 21, 21, 21, 55, 69, 55, 69, 69, 55, 69, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 91, 91, 91, 91, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 10, 11, 91, 91, 75,
	79, 79, 79, 75, 79, 79, 79, 75, 79, 79, 79, 75, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 1,
	1, 10, 11, 33, 33, 26, 27, 55, 69, 55, 69, 69, 55, 69, 55, 92,
	92, 92, 92, 93, 93, 93, 93, 75, 76, 1, 62, 75, 76, 33, 63, 92,
	92, 92, 92, 93, 93, 93, 93, 1, 1, 1, 1, 33, 33, 33, 33, 55,
	69, 55, 69, 69, 55, 69, 55, 92, 92, 92, 92, 93, 93, 93, 93, 92,
	92, 92, 92, 93, 93, 93, 93, 1, 1, 75, 76, 33, 33, 75, 76, 75,
	76, 55, 69, 75, 76, 69, 55, 92, 92, 92, 92, 93, 93, 93, 93, 55,
	69, 75, 76, 69, 55, 75, 76, 92, 92, 92, 92, 93, 93, 93, 93, 75,
	76, 55, 69, 75, 76, 69, 55, 75, 76, 55, 69, 75, 76, 69, 55, 55,
	69, 75, 76, 69, 55, 75, 76, 55, 69, 75, 76, 69, 55, 75, 76, 4,
	4, 4, 4, 20, 20, 20, 20, 55, 69, 55, 69, 69, 55, 69, 55, 92,
	92, 92, 92, 93, 93, 93, 93, 40, 40, 40, 40, 56, 56, 56, 56, 55,
	69, 1, 1, 69, 55, 33, 33, 55, 69, 55, 69, 69, 55, 69, 55,
}
