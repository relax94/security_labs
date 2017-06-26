var bigInt = require("big-integer");

const extendedEuklid = (fi, e) => {
    console.log('START COMPUTE EXTENDED EUKLIDIAN');
    const table = [[fi, e], [bigInt(-1), bigInt(-1)], [bigInt(1), bigInt(0)], [bigInt(0), bigInt(1)]];
    let r = bigInt(0);
    let i = bigInt(2);
    do {
        r = table[0][i - 2].mod(table[0][i - 1]);
        const q = table[0][i - 2].divide(table[0][i - 1]);
        const l = table[2][i - 2] - (q * table[2][i - 1])
        const b = table[3][i - 2] - (q * table[3][i - 1])

        table[0][i] = r;
        table[1][i] = q;
        table[2][i] = l;
        table[3][i] = b;
        i++;
    } while (!r.equals(1));

    const d = table[3][--i];
    return bigInt(d).greater(0) ? d : d + fi;
};


const computeLeft = (x, B, p, e) => {
    console.log('START FOUND AN ANSWER')
    for (let i = 1; i < x; i++) {
        const c = bigInt(e).pow(i);
        const l = bigInt(B).multiply(c);
        const r = l.mod(p);
        if (right[r]) {
            console.log('ANSWER R = ', r);
            console.log('ANSWER X = ', i);
            return;
        }
    }
};

const right = {};
const computeRight = (x, B, p, g) => {
    console.log('START COMPUTE RIGHT PART')
    for (let i = 1; i < x; i++) {
        const c = bigInt(g).pow(B).pow(i);
        const r = c.mod(p);
        right[r] = i;
    }
};

const computeB = (p) => {
    console.log('START COMPUTE B');
    let i = 0;
    while (bigInt(2).pow(i).lesser(p)) {
        i++;
    }
    const bresponse = bigInt(2).pow(bigInt(i).divide(2));
    console.log('FOUND B ', bresponse);
    return bresponse;
};

const g = bigInt('7');
const p = bigInt('683');
const h = bigInt('510');

const B = computeB(p);
const e = extendedEuklid(p, g);

//console.log(extendedEuklid(p, g))

//const left = bigInt(7).modPow(B, p);
// computeLeft(32, B, p, e);
computeRight(B, B, p, g);
computeLeft(B, h, p, e);
