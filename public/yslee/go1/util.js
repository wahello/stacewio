//document
document.addEventListener("keydown", OneDownHandler, false);
document.addEventListener("keyup", OneUpHandler, false);
document.addEventListener("touchstart", OneDownHandler, false);
document.addEventListener("touchend", OneUpHandler, false);
document.addEventListener("mousedown", OneDownHandler, false);
document.addEventListener("mouseup", OneUpHandler);
function pressKey() {
    sio.emit('cPressKey');
}
let onePressed = false;
function OneDownHandler(e) {
    if (onePressed == true) {
        return;
    }
    if (e.key == "Spacebar" || e.key == " " ||
        e.type == "mousedown" || e.type == "touchstart") {
        onePressed = true;
        pressKey();
    }
}
function OneUpHandler(e) {
    if (e.key == "Spacebar" || e.key == " " ||
        e.type == "mouseup" || e.type == "touchend") {
        onePressed = false;
    }
}

class myCanvas {
    static getCXH(sX) { return sX * canvas.width / sCanvasXY; }
    static getCY(sY) { return canvas.height - (sY * canvas.height / sCanvasXY); }
    static getCR(sR) { return sR * canvas.height / sCanvasXY; }
    static DrawCircle(sX, sY, sR, color) {
        ctx.beginPath();
        ctx.arc(myCanvas.getCXH(sX), myCanvas.getCY(sY), myCanvas.getCR(sR), 0, Math.PI * 10);
        ctx.fillStyle = color;
        ctx.fill();
        ctx.closePath();
    }
    static DrawUser(sX, sY, sW, sH, img) {
        let cX = myCanvas.getCXH(sX);
        let cY = myCanvas.getCY(sY);
        let cW = myCanvas.getCXH(sW);
        let cH = myCanvas.getCXH(sH);
        ctx.drawImage(img, cX - cW, cY - cH, cW + cW, cH);//x,y,w,h
    }
}