function pressKey() {
    sio.emit('cPressKey');
}
sio.on('sFrame', function (msg) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    drawMsgDecoder(msg);
});

