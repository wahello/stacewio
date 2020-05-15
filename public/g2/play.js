let playerPosX = 0;
let playerHeight = 50;

function drawPlayer() {
    ctx.beginPath();
    //ctx.rect(paddleX, canvas.height - paddleHeight, paddleWidth, paddleHeight);
    ctx.rect(0 + playerPosX, canvas.height, 10, -playerHeight);//x,y,w,h
    ctx.fillStyle = "#0095DD";
    ctx.fill();
    ctx.closePath();
}

function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    drawPlayer();
}


sio.on('sFrame', function (msg) {
    playerPosX++;
    draw();
});