const lexicalAlpha = () => {
    // 1  uniq = set(args)
    //2 sort
    //compare, 2 loop, uniq & args loop, if uniqLetter== argsLetter, acc+=letter

    //2 sloution, use Map, alpha := map[letter: count],  2 sort By key, 3 accum loop, by map, value

    //3 sort letters, O(n)
    return args.split("").sort().join("") //split string -> sort, toString()

}