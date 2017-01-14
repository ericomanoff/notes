const decrypt = flag =>
  flag.split('')
      .map((ch, i) => String.fromCharCode(ch.charCodeAt(0) - i))
      .join('')

let f = 'c8e49=<@';

console.log(`decrypt('${f}') -> ${decrypt(f)}`)
