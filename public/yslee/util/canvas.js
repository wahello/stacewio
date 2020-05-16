const canvas = document.getElementById("canvas");
document.addEventListener("contextmenu", handleContextualMenu);
function handleContextualMenu(event) {
    event.preventDefault();
}
const ctx = canvas.getContext("2d");
const sCanvasXY = 1000;
