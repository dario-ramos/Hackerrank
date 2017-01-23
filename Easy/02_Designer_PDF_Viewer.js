process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();    
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

function main() {
    h = readLine().split(' ');
    h = h.map(Number);
    var word = readLine().split('');
    var maxLetterHeight = word.reduce(function(maxHeight, letter){
        var letterIndex = letter.charCodeAt(0) - 'a'.charCodeAt(0);
        return (h[letterIndex] > maxHeight)? h[letterIndex] : maxHeight;
    }, 0);    
    console.log(maxLetterHeight * word.length);
}