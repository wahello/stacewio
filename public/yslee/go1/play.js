let playerPosX = 0;
let playerHeight = 50;

function drawBlock(msg, nFindIndex) {
    let sX, sY, sType;

    let multi = msgToken([msg, nFindIndex]);
    sX = parseInt(multi[0]);

    multi = msgToken([msg, multi[1]]);
    sY = parseInt(multi[0]);

    multi = msgToken([msg, multi[1]]);
    sType = parseInt(multi[0]);

    myCanvas.DrawCircle(sX, sY, 20, "sienna");

    return nFindIndex;
}

function msgToken(multi) {
    let nStart = multi[1] + 1;
    multi[1] = multi[0].indexOf(",", nStart);
    multi[0] = multi[0].substring(nStart, multi[1]);
    return [multi[0], multi[1]];
}

function drawUser(msg, nFindIndex) {
    let sC, sX, sW, sH, sS, sP;
    let multi = msgToken([msg, nFindIndex]);
    sC = parseInt(multi[0]);
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

    if (sS < 0) {
        img = sP % 20 < 10 ? imgl1 : imgl2

    } else {
        img = sP % 20 < 10 ? imgr1 : imgr2
    }
    myCanvas.DrawUser(sX, 0, 30, 60, img);
    return nFindIndex
}
//닉네임 어떻게 할 지...
//180프레임간 무적


function drawMetaInfo(msg, nFindIndex) {
    //이런 일은 하지 말아야 했는데, 난 그 사실을 몰랐어.
    //이제 와서 후회한들 뭐하리 난 바보가 돼 버린걸.
}

function drawMsgDecoder(msg) {
    let nFindIndex = 0;
    while ((nFindIndex = msg.indexOf(".", nFindIndex)) > -1) {
        let charType = msg.charAt(nFindIndex + 1);
        switch (charType) {
            case "b":
                nFindIndex = drawBlock(msg, nFindIndex + 1);
                break;
            case "u":
                nFindIndex = drawUser(msg, nFindIndex + 1);
                break;
            case "m":
                nFindIndex = drawMetaInfo(msg, nFindIndex + 1);
            default:
                break;
        }
    }
}


sio.on('sFrame', function (msg) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    drawMsgDecoder(msg);
});