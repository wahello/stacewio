//class Draw{ IE Class 지원 안 함..
function getCXH(sX) { return sX * canvas.width / sCanvasXY; }
function getCY(sY) { return canvas.height - (sY * canvas.height / sCanvasXY); }
function getCR(sR) { return sR * canvas.height / sCanvasXY; }
function DrawCircle(sX, sY, sR, color) {
    ctx.beginPath();
    ctx.arc(getCXH(sX), getCY(sY), getCR(sR), 0, Math.PI * 10);
    ctx.fillStyle = color;
    ctx.fill();
    ctx.closePath();
}
function DrawUserSquare(sX, sY, sW, sH, color) {
    let cX = getCXH(sX);
    let cY = getCY(sY);
    let cW = getCXH(sW);
    let cH = getCXH(sH);
    ctx.beginPath();
    ctx.rect(cX - cW, cY - cH, cW + cW, cH);
    ctx.fillStyle = color;
    ctx.fill();
    ctx.closePath();
}
function DrawUser(sX, sY, sW, sH, img) {
    let cX = getCXH(sX);
    let cY = getCY(sY);
    let cW = getCXH(sW);
    let cH = getCXH(sH);
    ctx.drawImage(img, cX - cW, cY - cH, cW + cW, cH);//x,y,w,h
}
//}IE Class 지원 안 함..

function decodeAndDrawBlock(msg, nFindIndex) {
    let sX, sY, sType;

    let multi = msgToken([msg, nFindIndex]);
    sX = parseInt(multi[0]);

    multi = msgToken([msg, multi[1]]);
    sY = parseInt(multi[0]);

    multi = msgToken([msg, multi[1]]);
    sType = parseInt(multi[0]);

    DrawCircle(sX, sY, 20, "sienna");

    return multi[1];
}


function decodeUser(msg, nFindIndex) {
    let sC, sX, sW, sH, sS, sP;
    let multi = msgToken([msg, nFindIndex]);
    sC = multi[0];
    multi = msgToken([msg, multi[1]]);
    sX = parseInt(multi[0]);
    multi = msgToken([msg, multi[1]]);
    sW = parseInt(multi[0]);
    multi = msgToken([msg, multi[1]]);
    sH = parseInt(multi[0]);
    multi = msgToken([msg, multi[1]]);
    sS = parseInt(multi[0]);
    multi = msgToken([msg, multi[1]]);
    sP = parseInt(multi[0]);
    if (sC == sio.io.engine.id) {
        meX = sX; meW = sW; meH = sH; meS = sS; meP = sP;
    }
    else {
        drawUserInfo(sX, sW, sH, sS, sP, false);
    }

    return multi[1];
}
function drawUserInfo(sX, sW, sH, sS, sP, me) {
    if (sS < 0) {
        img = sP % 20 < 10 ? imgl1 : imgl2;
    } else {
        img = sP % 20 < 10 ? imgr1 : imgr2;
    }

    if (me && sP < 150) {
        console.log(me);
        DrawCircle(sX, 0, 150 - sP, "black");
    }

    DrawUser(sX, 0, sW, sH, img);
}

let meX, meW, meH, meS, meP;
function drawMe() {
    drawUserInfo(meX, meW, meH, meS, meP, true);
}

function drawMetaInfo(msg, nFindIndex) {

}
