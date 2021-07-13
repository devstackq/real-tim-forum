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

const find2Max = (arr) => {
  let max = 0;
  let second = 0;
  for (let i = 0; i < arr.length; i++) {
    //4,9,16,2

    if (max < arr[i]) {
      second = max;
      max = arr[i];
      console.log(max, second);
    }
    //else if arr[i]> second{second = arr[i]}
    if (second < max && second < arr[i] && arr[i] != max) {
      second = arr[i];
      console.log(max, second);
    }
  }
  return [max, second];
};

const findMinEven = (arr) => {
  //else -1
  //loop [5,8,11,9,10,20,3,7]
  let min = 0;

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] % 2 == 1) {
      if (min > arr[i]) {
        min = arr[i]; // -1 > 1 f,
      } else if (min == 0) {
        min = arr[i];
      }
    }
  }
  if (min == 0) {
    return -1;
  }
  return min;
};
