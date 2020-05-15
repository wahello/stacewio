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
    if (e.key == "Spacebar" || e.key == " ") {
        onePressed = true;
        pressKey();
    }
}
function OneUpHandler(e) {
    if (e.key == "Spacebar" || e.key == " ") {
        onePressed = false;
    }
}