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
    var a_temp = readLine().split(' ');
    var a = parseInt(a_temp[0]);
    var b = parseInt(a_temp[1]);
    var c = parseInt(a_temp[2]);
    var d = parseInt(a_temp[3]);
    var e = parseInt(a_temp[4]);

    var sum = a;
    var min = a;
    var max = a;
    //Compute sum, min and max values
    for(var i=1; i<5; i++){
        var item = parseInt(a_temp[i]);
        sum += item;
        if( item < min ){
            min = item;
        }
        if( item > max ){
            max = item;
        }
    }
    //Min sum is sum-max, Max sum is sum-min
    console.log(sum-max, sum-min);
}
