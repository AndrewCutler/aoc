import * as fs from 'fs';
import * as readline from 'readline';

// day 1
async function day1() {
	const stream = fs.createReadStream('../data/day1.txt');

	const rl = readline.createInterface({
		input: stream,
		crlfDelay: Infinity
	});

	const elves: string[][] = [];
	let elf: string[] = [];
	for await (const line of rl) {
		if (line === '') {
			elves.push(elf);
			elf = [];
			continue;
		}
		elf.push(line);
	}
	const sums = elves.map((elf) =>
		elf.reduce((prev, curr) => {
			prev += parseInt(curr, 10);
			return prev;
		}, 0)
	);

	console.log({ sums });
	const [one, two, three] = sums.sort((a, b) => a - b).reverse();
	// console.log({ one, two, three });

	rl.close();
	console.log({
		partOne: one,
		partTwo: one + two + three
	});
}

async function day2() {
	const outcome = {
		Win: 6,
		Lose: 0,
		Draw: 3
	} as const;
	const shape = {
		A: {
			type: 'rock',
			beats: 'scissors',
			losesTo: 'paper',
			outcome: undefined,
			value: 1
		},
		X: {
			type: 'rock',
			beats: 'scissors',
			losesTo: 'paper',
			outcome: outcome.Lose,
			value: 1
		},
		B: {
			type: 'paper',
			beats: 'rock',
			losesTo: 'scissors',
			outcome: undefined,
			value: 2
		},
		Y: {
			type: 'paper',
			beats: 'rock',
			losesTo: 'scissors',
			outcome: outcome.Draw,
			value: 2
		},
		C: {
			type: 'scissors',
			beats: 'paper',
			losesTo: 'rock',
			value: 3,
			outcome: undefined
		},
		Z: {
			type: 'scissors',
			beats: 'paper',
			losesTo: 'rock',
			outcome: outcome.Win,
			value: 3
		}
	} as const;

	function getShape(...args: string[]) {
		const shapes = [];
		for (let i = 0; i < args.length; i++) {
			const curr = shape[args[i] as keyof typeof shape];
			shapes.push(curr);
		}

		return shapes;
	}

	const stream = fs.createReadStream('../data/day2.txt');

	const rl = readline.createInterface({
		input: stream,
		crlfDelay: Infinity
	});

	let partOneSum = 0;
	let partTwoSum = 0;

	for await (const line of rl) {
		const [them, us] = getShape(...line.split(' '));

		// part one
		if (us.beats === them.type) {
			partOneSum += outcome.Win + us.value;
		}
		if (us.type === them.type) {
			partOneSum += outcome.Draw + us.value;
		}
		if (us.losesTo === them.type) {
			partOneSum += outcome.Lose + us.value;
		}

		// part two
		if (us.outcome === outcome.Win) {
			partTwoSum +=
				outcome.Win +
				Object.values(shape).find((s) => s.beats === them.type)!.value;
		}
		if (us.outcome === outcome.Draw) {
			partTwoSum += outcome.Draw + them.value;
		}
		if (us.outcome === outcome.Lose) {
			partTwoSum +=
				outcome.Lose +
				Object.values(shape).find((s) => s.losesTo === them.type)!
					.value;
		}
	}

	rl.close();
	console.log({ partOne: partOneSum, partTwo: partTwoSum });
}

async function day3() {
	const stream = fs.createReadStream('../data/day3.txt');

	const rl = readline.createInterface({
		input: stream,
		crlfDelay: Infinity
	});

	function getCharValue(char: string): number {
		const letters = [
			...'abcdefghijklmnopqrstuvwxyz'.split(''),
			...'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('')
		];

		return letters.indexOf(char) + 1;
	}

	let partOneSum = 0;
	let partTwoIteration: string[] = [];
	let partTwoSum = 0;
	let i = 0;
	for await (const line of rl) {
		if (i === 3) {
			const [longest, two, three] = partTwoIteration;

			for (const char of longest.split('')) {
				// char is in longest, two and three
				if (three.includes(char) && two.includes(char)) {
					partTwoSum += getCharValue(char);
				}
			}

			(partTwoIteration = [line]), (i = 0);
		} else {
			partTwoIteration = [...partTwoIteration, line].sort(
				(a, b) => a.length - b.length
			);

			++i;
		}

		const chars = line.split('');
		const midpoint = chars.length / 2;
		const left = chars.slice(0, midpoint);
		const right = chars.slice(midpoint);

		for (const char of left) {
			if (right.includes(char)) {
				partOneSum += getCharValue(char);
				break;
			}
		}
	}

	console.log({ partOne: partOneSum, partTwo: partTwoSum /* wrong */ });
}

async function day4() {
	const stream = fs.createReadStream('../data/day4.txt');
	const rl = readline.createInterface({
		input: stream,
		crlfDelay: Infinity
	});

	let partOneCount = 0;
	let partTwoCount = 0;
	for await (const line of rl) {
		const [first, second]: number[][] = line
			.split(',')
			.map((s) => s.split('-').map((n) => parseInt(n, 10)));

		// part one
		if (
			(first[0] <= second[0] && first[1] >= second[1]) ||
			(second[0] <= first[0] && second[1] >= first[1])
		)
			++partOneCount;

		// part two
		if (
			(second[0] >= first[0] && second[0] <= first[1]) ||
			(second[1] >= first[0] && second[1] <= first[1]) || 
			(first[0] >= second[0] && first[0] <= second[1]) ||
			(first[1] >= second[0] && first[1] <= second[1])
		)
			++partTwoCount;
	}

	console.log({ partOne: partOneCount, partTwo: partTwoCount });
}

(async () => {
	// await day1();
	// await day2();
	// await day3();
	await day4();
})();
