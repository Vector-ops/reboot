let newAdder = fn(x) {
  fn (y) {
    return x + y;
  }
}

let addTwo = newAdder(2);

addTwo(3);

if (5 == 6) {
  return true;
} else {
  return false;
}