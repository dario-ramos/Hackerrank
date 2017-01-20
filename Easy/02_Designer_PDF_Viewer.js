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
    var word = readLine();
    var maxLetterHeight = 0;
    for(var i=0; i<word.length; i++){
        var letterIndex = word.charCodeAt(i)-'a'.charCodeAt(0);
        if(h[letterIndex] > maxLetterHeight){
            maxLetterHeight = h[letterIndex];
        }
    }
    console.log(maxLetterHeight * word.length);
}