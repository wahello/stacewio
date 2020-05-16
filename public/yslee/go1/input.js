function pressKey(e) {
    sio.emit('cPressKey');
}

//Key input
document.addEventListener("mousedown", pressKey, false);
document.addEventListener("touchstart", pressKey, false);
document.addEventListener("keydown", KeyDownHandler, false);
document.addEventListener("keyup", KeyUpHandler, false);
let spaceKey;
function KeyDownHandler(e) {
    if (spaceKey == true)
        return;
    if (e.key == "Spacebar" || e.key == " ") {
        pressKey(e);
        spaceKey = true;
    }
}
function KeyUpHandler(e) {
    if (e.key == "Spacebar" || e.key == " ")
        spaceKey = false;
}




