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
