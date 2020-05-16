function msgToken(multi) {
    let nStart = multi[1] + 1;
    multi[1] = multi[0].indexOf(",", nStart);
    multi[0] = multi[0].substring(nStart, multi[1]);
    return [multi[0], multi[1]];
}

function drawMsgDecoder(msg) {
    let nFindIndex = 0;
    while ((nFindIndex = msg.indexOf(".", nFindIndex)) > -1) {
        let charType = msg.charAt(nFindIndex + 1);
        switch (charType) {
            case "b":
                nFindIndex = decodeAndDrawBlock(msg, nFindIndex + 1);
                break;
            case "u":
                nFindIndex = decodeUser(msg, nFindIndex + 1);
                break;
            case "m":
                nFindIndex = drawMetaInfo(msg, nFindIndex + 1);
            default:
                break;
        }
    }

    drawMe();
}