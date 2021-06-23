const firstInNum = (arr, target) => {
  //find first, target in array
  let n = -1;

  for (let i = 0; i < arr.length; i++) {
    if (n == -1 && arr[i] === target) {
      n = i;
      break;
    }
  }
  return n;
};

const findMax = (seq) => {
  let first = seq[0];

  for (let i = 0; i < seq.length; i++) {
    if (first < seq[i]) {
      first = seq[i];
    }
  }
  return first;
};

//find and sort by alpha
const lexGraph = (seq) => {
  let res = "";
  let uniqLetters = set(seq);
  //sort by alpha, find uniq, then count, same letter
  //use map or set -> set, count each word
  console.log(uniqLetters);
  // for(let i =0;i < seq.length; i++) {
  //   if(seq[i]  )
  // }
};

const set = (value) => {
  //return uniq value in seq
  let temp = value[0];
//   'qeroponvlkqsqwer'
    //curerentIndexValue != nextIndexValue
//index "askalskqk" - 0,3, / 
//2 loop, 0index, searchSameLetter, recursive rewrite strinmg ?
// 0,3 -> 
let accum = ""
for (let i = 0; i < seq.length; i++) {
    if(q != )
    cIv != nIv {
        temp += seq[i]
    }
  }

};
