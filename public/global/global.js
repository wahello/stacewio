const canvas = document.getElementById("canvas");
document.addEventListener("contextmenu", handleContextualMenu);
function handleContextualMenu(event) {
    event.preventDefault();
}
const ctx = canvas.getContext("2d");
const sCanvasXY = 1000;

class myCanvas {
    static getCX(sX) { return sX * canvas.width / sCanvasXY; }
    static getCY(sY) { return canvas.height - (sY * canvas.height / sCanvasXY); }
    static getCR(sR) { return sR * canvas.height / sCanvasXY; }
    static DrawCircle(sX, sY, sR, color) {
        ctx.beginPath();
        ctx.arc(myCanvas.getCX(sX), myCanvas.getCY(sY), myCanvas.getCR(sR), 0, Math.PI * 10);
        ctx.fillStyle = color;
        ctx.fill();
        ctx.closePath();
    }
    static DrawUser(sX, sH, sW, color) {
        let cX = myCanvas.getCX(sX);
        let cH = myCanvas.getCY(sH);
        let cW = myCanvas.getCX(sW);
        ctx.beginPath();
        ctx.rect(cX - cW, cH, cW + cW, cH);//x,y,w,h
        ctx.fillStyle = color;
        ctx.fill();
        ctx.closePath();
    }
}



// class Block {
//     constructor(index, hash, preHash, data, timeStamp) {
//         this.index = index;
//         this.hash = hash;
//         this.preHash = preHash;
//         this.data = data;
//         this.timeStamp = timeStamp;
//     }
//     static calculateBlockHash(index, preHash, timeStamp, data) {
//         return CryptoJS.SHA256(index + preHash + timeStamp + data).toString();
//     }
//     static validateStructure(aBlock) {
//         return typeof aBlock.index === "number"
//             && typeof aBlock.hash === "string"
//             && typeof aBlock.preHash === "string"
//             && typeof aBlock.timeStamp === "number"
//             && typeof aBlock.data === "string";
//     }
// }
// const genesisBlock = new Block(0, "202020202020202", "", "Hello World", 123456);