"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var __asyncValues = (this && this.__asyncValues) || function (o) {
    if (!Symbol.asyncIterator) throw new TypeError("Symbol.asyncIterator is not defined.");
    var m = o[Symbol.asyncIterator], i;
    return m ? m.call(o) : (o = typeof __values === "function" ? __values(o) : o[Symbol.iterator](), i = {}, verb("next"), verb("throw"), verb("return"), i[Symbol.asyncIterator] = function () { return this; }, i);
    function verb(n) { i[n] = o[n] && function (v) { return new Promise(function (resolve, reject) { v = o[n](v), settle(resolve, reject, v.done, v.value); }); }; }
    function settle(resolve, reject, d, v) { Promise.resolve(v).then(function(v) { resolve({ value: v, done: d }); }, reject); }
};
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
Object.defineProperty(exports, "__esModule", { value: true });
var fs = require("fs");
var readline = require("readline");
// day 1
function day1() {
    var _a, e_1, _b, _c;
    return __awaiter(this, void 0, void 0, function () {
        var stream, rl, elves, elf, _d, rl_1, rl_1_1, line, e_1_1, sums, _e, one, two, three;
        return __generator(this, function (_f) {
            switch (_f.label) {
                case 0:
                    stream = fs.createReadStream('./data/day1.txt');
                    rl = readline.createInterface({
                        input: stream,
                        crlfDelay: Infinity
                    });
                    elves = [];
                    elf = [];
                    _f.label = 1;
                case 1:
                    _f.trys.push([1, 6, 7, 12]);
                    _d = true, rl_1 = __asyncValues(rl);
                    _f.label = 2;
                case 2: return [4 /*yield*/, rl_1.next()];
                case 3:
                    if (!(rl_1_1 = _f.sent(), _a = rl_1_1.done, !_a)) return [3 /*break*/, 5];
                    _c = rl_1_1.value;
                    _d = false;
                    line = _c;
                    if (line === '') {
                        elves.push(elf);
                        elf = [];
                        return [3 /*break*/, 4];
                    }
                    elf.push(line);
                    _f.label = 4;
                case 4:
                    _d = true;
                    return [3 /*break*/, 2];
                case 5: return [3 /*break*/, 12];
                case 6:
                    e_1_1 = _f.sent();
                    e_1 = { error: e_1_1 };
                    return [3 /*break*/, 12];
                case 7:
                    _f.trys.push([7, , 10, 11]);
                    if (!(!_d && !_a && (_b = rl_1.return))) return [3 /*break*/, 9];
                    return [4 /*yield*/, _b.call(rl_1)];
                case 8:
                    _f.sent();
                    _f.label = 9;
                case 9: return [3 /*break*/, 11];
                case 10:
                    if (e_1) throw e_1.error;
                    return [7 /*endfinally*/];
                case 11: return [7 /*endfinally*/];
                case 12:
                    sums = elves.map(function (elf) {
                        return elf.reduce(function (prev, curr) {
                            prev += parseInt(curr, 10);
                            return prev;
                        }, 0);
                    });
                    console.log({ sums: sums });
                    _e = sums.sort(function (a, b) { return a - b; }).reverse(), one = _e[0], two = _e[1], three = _e[2];
                    // console.log({ one, two, three });
                    rl.close();
                    console.log({
                        partOne: one,
                        partTwo: one + two + three
                    });
                    return [2 /*return*/];
            }
        });
    });
}
function day2() {
    var _a, e_2, _b, _c;
    return __awaiter(this, void 0, void 0, function () {
        function getShape() {
            var args = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                args[_i] = arguments[_i];
            }
            var shapes = [];
            for (var i = 0; i < args.length; i++) {
                var curr = shape[args[i]];
                shapes.push(curr);
            }
            return shapes;
        }
        var outcome, shape, stream, rl, partOneSum, partTwoSum, _loop_1, _d, rl_2, rl_2_1, e_2_1;
        return __generator(this, function (_e) {
            switch (_e.label) {
                case 0:
                    outcome = {
                        Win: 6,
                        Lose: 0,
                        Draw: 3
                    };
                    shape = {
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
                    };
                    stream = fs.createReadStream('./data/day2.txt');
                    rl = readline.createInterface({
                        input: stream,
                        crlfDelay: Infinity
                    });
                    partOneSum = 0;
                    partTwoSum = 0;
                    _e.label = 1;
                case 1:
                    _e.trys.push([1, 6, 7, 12]);
                    _loop_1 = function () {
                        _c = rl_2_1.value;
                        _d = false;
                        var line = _c;
                        var _f = getShape.apply(void 0, line.split(' ')), them = _f[0], us = _f[1];
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
                                    Object.values(shape).find(function (s) { return s.beats === them.type; }).value;
                        }
                        if (us.outcome === outcome.Draw) {
                            partTwoSum += outcome.Draw + them.value;
                        }
                        if (us.outcome === outcome.Lose) {
                            partTwoSum +=
                                outcome.Lose +
                                    Object.values(shape).find(function (s) { return s.losesTo === them.type; })
                                        .value;
                        }
                    };
                    _d = true, rl_2 = __asyncValues(rl);
                    _e.label = 2;
                case 2: return [4 /*yield*/, rl_2.next()];
                case 3:
                    if (!(rl_2_1 = _e.sent(), _a = rl_2_1.done, !_a)) return [3 /*break*/, 5];
                    _loop_1();
                    _e.label = 4;
                case 4:
                    _d = true;
                    return [3 /*break*/, 2];
                case 5: return [3 /*break*/, 12];
                case 6:
                    e_2_1 = _e.sent();
                    e_2 = { error: e_2_1 };
                    return [3 /*break*/, 12];
                case 7:
                    _e.trys.push([7, , 10, 11]);
                    if (!(!_d && !_a && (_b = rl_2.return))) return [3 /*break*/, 9];
                    return [4 /*yield*/, _b.call(rl_2)];
                case 8:
                    _e.sent();
                    _e.label = 9;
                case 9: return [3 /*break*/, 11];
                case 10:
                    if (e_2) throw e_2.error;
                    return [7 /*endfinally*/];
                case 11: return [7 /*endfinally*/];
                case 12:
                    rl.close();
                    console.log({ partOne: partOneSum, partTwo: partTwoSum });
                    return [2 /*return*/];
            }
        });
    });
}
function day3() {
    var _a, e_3, _b, _c;
    return __awaiter(this, void 0, void 0, function () {
        function getCharValue(char) {
            var letters = __spreadArray(__spreadArray([], 'abcdefghijklmnopqrstuvwxyz'.split(''), true), 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split(''), true);
            return letters.indexOf(char) + 1;
        }
        var stream, rl, partOneSum, partTwoIteration, partTwoSum, i, _d, rl_3, rl_3_1, line, longest, two, three, _i, _e, char, chars, midpoint, left, right, _f, left_1, char, e_3_1;
        return __generator(this, function (_g) {
            switch (_g.label) {
                case 0:
                    stream = fs.createReadStream('./data/day3.txt');
                    rl = readline.createInterface({
                        input: stream,
                        crlfDelay: Infinity
                    });
                    partOneSum = 0;
                    partTwoIteration = [];
                    partTwoSum = 0;
                    i = 0;
                    _g.label = 1;
                case 1:
                    _g.trys.push([1, 6, 7, 12]);
                    _d = true, rl_3 = __asyncValues(rl);
                    _g.label = 2;
                case 2: return [4 /*yield*/, rl_3.next()];
                case 3:
                    if (!(rl_3_1 = _g.sent(), _a = rl_3_1.done, !_a)) return [3 /*break*/, 5];
                    _c = rl_3_1.value;
                    _d = false;
                    line = _c;
                    if (i === 3) {
                        longest = partTwoIteration[0], two = partTwoIteration[1], three = partTwoIteration[2];
                        for (_i = 0, _e = longest.split(''); _i < _e.length; _i++) {
                            char = _e[_i];
                            // char is in longest, two and three
                            if (three.includes(char) && two.includes(char)) {
                                partTwoSum += getCharValue(char);
                            }
                        }
                        (partTwoIteration = [line]), (i = 0);
                    }
                    else {
                        partTwoIteration = __spreadArray(__spreadArray([], partTwoIteration, true), [line], false).sort(function (a, b) { return a.length - b.length; });
                        ++i;
                    }
                    chars = line.split('');
                    midpoint = chars.length / 2;
                    left = chars.slice(0, midpoint);
                    right = chars.slice(midpoint);
                    for (_f = 0, left_1 = left; _f < left_1.length; _f++) {
                        char = left_1[_f];
                        if (right.includes(char)) {
                            partOneSum += getCharValue(char);
                            break;
                        }
                    }
                    _g.label = 4;
                case 4:
                    _d = true;
                    return [3 /*break*/, 2];
                case 5: return [3 /*break*/, 12];
                case 6:
                    e_3_1 = _g.sent();
                    e_3 = { error: e_3_1 };
                    return [3 /*break*/, 12];
                case 7:
                    _g.trys.push([7, , 10, 11]);
                    if (!(!_d && !_a && (_b = rl_3.return))) return [3 /*break*/, 9];
                    return [4 /*yield*/, _b.call(rl_3)];
                case 8:
                    _g.sent();
                    _g.label = 9;
                case 9: return [3 /*break*/, 11];
                case 10:
                    if (e_3) throw e_3.error;
                    return [7 /*endfinally*/];
                case 11: return [7 /*endfinally*/];
                case 12:
                    console.log({ partOne: partOneSum, partTwo: partTwoSum /* wrong */ });
                    return [2 /*return*/];
            }
        });
    });
}
function day4() {
    var _a, e_4, _b, _c;
    return __awaiter(this, void 0, void 0, function () {
        var stream, rl, partOneCount, partTwoCount, _d, rl_4, rl_4_1, line, _e, first, second, e_4_1;
        return __generator(this, function (_f) {
            switch (_f.label) {
                case 0:
                    stream = fs.createReadStream('./data/day4.txt');
                    rl = readline.createInterface({
                        input: stream,
                        crlfDelay: Infinity
                    });
                    partOneCount = 0;
                    partTwoCount = 0;
                    _f.label = 1;
                case 1:
                    _f.trys.push([1, 6, 7, 12]);
                    _d = true, rl_4 = __asyncValues(rl);
                    _f.label = 2;
                case 2: return [4 /*yield*/, rl_4.next()];
                case 3:
                    if (!(rl_4_1 = _f.sent(), _a = rl_4_1.done, !_a)) return [3 /*break*/, 5];
                    _c = rl_4_1.value;
                    _d = false;
                    line = _c;
                    _e = line
                        .split(',')
                        .map(function (s) { return s.split('-').map(function (n) { return parseInt(n, 10); }); }), first = _e[0], second = _e[1];
                    // part one
                    if ((first[0] <= second[0] && first[1] >= second[1]) ||
                        (second[0] <= first[0] && second[1] >= first[1]))
                        ++partOneCount;
                    // part two
                    if ((second[0] >= first[0] && second[0] <= first[1]) ||
                        (second[1] >= first[0] && second[1] <= first[1]) ||
                        (first[0] >= second[0] && first[0] <= second[1]) ||
                        (first[1] >= second[0] && first[1] <= second[1]))
                        ++partTwoCount;
                    _f.label = 4;
                case 4:
                    _d = true;
                    return [3 /*break*/, 2];
                case 5: return [3 /*break*/, 12];
                case 6:
                    e_4_1 = _f.sent();
                    e_4 = { error: e_4_1 };
                    return [3 /*break*/, 12];
                case 7:
                    _f.trys.push([7, , 10, 11]);
                    if (!(!_d && !_a && (_b = rl_4.return))) return [3 /*break*/, 9];
                    return [4 /*yield*/, _b.call(rl_4)];
                case 8:
                    _f.sent();
                    _f.label = 9;
                case 9: return [3 /*break*/, 11];
                case 10:
                    if (e_4) throw e_4.error;
                    return [7 /*endfinally*/];
                case 11: return [7 /*endfinally*/];
                case 12:
                    console.log({ partOne: partOneCount, partTwo: partTwoCount });
                    return [2 /*return*/];
            }
        });
    });
}
(function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: 
            // await day1();
            // await day2();
            // await day3();
            return [4 /*yield*/, day4()];
            case 1:
                // await day1();
                // await day2();
                // await day3();
                _a.sent();
                return [2 /*return*/];
        }
    });
}); })();
