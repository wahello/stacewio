document.addEventListener("keydown", OneDownHandler, false);
document.addEventListener("keyup", OneUpHandler, false);
document.addEventListener("touchstart", OneDownHandler, false);
document.addEventListener("touchend", OneUpHandler, false);
document.addEventListener("mousedown", OneDownHandler, false);
document.addEventListener("mouseup", OneUpHandler);
//document addEventListener.... 바꿔야됨 큰일날것 같은데..
//option???????????????????
function OneDownHandler(e) {
    if (onePressed == true) {
        return
    }
    if (e.key == "Spacebar" || e.key == " ") {
        onePressed = true;
    }
}
function OneUpHandler(e) {
    if (e.key == "Spacebar" || e.key == " ") {
        onePressed = false;
    }
}