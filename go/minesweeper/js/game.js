$(document).ready(init);

var clickCount = 0;

var ws = new WebSocket("ws://127.0.0.1:8080/ws");

ws.onmessage = function (event) {
    var data = JSON.parse(event.data);
    switch (data.name) {
        case "create":
            createTable(data.data)
            break;
        case "click":
            clickResult(data.data)
            break;
        case "flag":
            flagResult(data.data)
            break;
        case "check_around_flag":
            checkFlagResult(data.data)
            break;
    }
};

// 網頁初始化
function init() {
    $('#start').click(function () {
        clickCount = 0;
        $('#showTable').html("Loading...");
        setTable();
    });

    $('#primary').click(level);
    $('#middle').click(level);
    $('#high').click(level);
}

// 各級別遊戲設定
function level() {
    if ($(this).text() == "初級") {
        $("#row").val(9);
        $("#column").val(9);
        $("#m").val(10);
    }
    if ($(this).text() == "中級") {
        $("#row").val(16);
        $("#column").val(16);
        $("#m").val(40);
    }
    if ($(this).text() == "高級") {
        $("#row").val(16);
        $("#column").val(30);
        $("#m").val(99);
    }
}

// 建立遊戲地圖
function createTable(data) {
    $('#showTable').html(data);
    $('#showTable td').click(clickMap);
    $('#showTable td').dblclick(dbClickMap);
    $('#showTable td').hover(onMap, outMap);
    $('#showTable td').mousedown(function (event) {
        if ($(this).find('#content').is(':hidden')) {
            switch (event.which) {
                case 3:
                    flagGrip(this.id)
                    break;
            }
        }
    });
}

// 地圖點擊結果
function clickResult(data) {
    $.each(data.open, function (index, grip) {
        nowGrip = $('#' + grip.self)
        nowGrip.find('#content').text(grip.content)
        nowGrip.removeClass();
        print(grip.status, nowGrip)

    })

    if (data.result == "over") {
        isGameOver()
    }
    if (data.result == "clear") {
        gameClear()
    }
}

// 插旗結果
function flagResult(id) {
    point = $('#' + id)
    point.find('#flag').toggle();
}

// 檢查四周插旗數量，直接開啟四周格子
function checkFlagResult(data) {
    $.each(data.open, function (index, grip) {
        nowGrip = $('#' + grip.self)
        nowGrip.find('#content').text(grip.content)
        nowGrip.removeClass();
        print(grip.status, nowGrip)
    })

    if (data.result == "over") {
        isGameOver()
    }
    if (data.result == "clear") {
        gameClear()
    }
}

// 初次建立地圖
function setTable() {
    param = $("#row").val() + "," + $("#column").val() + "," + $("#m").val()
    ws.send(JSON.stringify({
        "name": "create",
        "param": param
    }));
}

// 地圖重置
function resetTable(trIndex, tdIndex) {
    $('#showTable tr').eq(trIndex).find('td').eq(tdIndex).click();
    param = $("#row").val() + "," + $("#column").val() + "," + $("#m").val()
    ws.send(param);
}

// 點擊地圖
function clickMap() {
    if ($(this).find('#flag').is(':visible')) {
        return;
    }

    // 已開啟的格子顯示提示
    if ($(this).find('#content').is(':visible')) {
        openedShow($(this))
        return;
    }

    ws.send(JSON.stringify({
        "name": "click",
        "param": this.id
    }));

    return;
}

// 格子插旗
function flagGrip(id) {
    ws.send(JSON.stringify({
        "name": "flag",
        "param": id
    }));
}

// 檢查四周的旗子 決定是否開啟四周個格子
function checkAroundFlag(id) {
    ws.send(JSON.stringify({
        "name": "check_around_flag",
        "param": id
    }));
}

// 雙擊地圖
function dbClickMap() {
    if ($(this).find('#content').is(':hidden')) {
        return;
    }

    var pNumber = $(this).find('#content').text().trim();
    var trIndex = $(this).closest('tr').index();
    var tdIndex = $(this).closest('td').index();
    var check = checkFlag(trIndex, tdIndex);

    if (pNumber != check.flag) {
        $.each(check.point, function () {
            if ($(this).find('#content').is(':hidden')) {
                $(this).addClass('checkArount');
            }
        });
        setTimeout(function () {
            $.each(check.point, function () {
                $(this).removeClass();
            });
        }, 250);
    } else {
        // 發 ws 檢查四周的旗子
        checkAroundFlag(this.id)
    }
}

// 檢查四周個旗子數
function checkFlag(trIndex, tdIndex) {
    var point = [];
    var flag = 0;

    // 判斷是否為最上方
    if (trIndex != 0) {
        // 上
        p = $('#showTable tr').eq(trIndex - 1).find('td').eq(tdIndex);
        if (p.find('#flag').is(':visible')) {
            flag++;
        } else {
            point.push(p);
        }

        // 判斷是否為最左側
        if (tdIndex != 0) {
            // 左上
            p = $('#showTable tr').eq(trIndex - 1).find('td').eq(tdIndex - 1);
            if (p.find('#flag').is(':visible')) {
                flag++;
            } else {
                point.push(p);
            }
        }

        // 右上
        p = $('#showTable tr').eq(trIndex - 1).find('td').eq(tdIndex + 1);
        if (p.find('#flag').is(':visible')) {
            flag++;
        } else {
            point.push(p);
        }
    }

    // 判斷是否為最左側
    if (tdIndex != 0) {
        // 左
        p = $('#showTable tr').eq(trIndex).find('td').eq(tdIndex - 1);
        if (p.find('#flag').is(':visible')) {
            flag++;
        } else {
            point.push(p);
        }
        // 左下
        p = $('#showTable tr').eq(trIndex + 1).find('td').eq(tdIndex - 1);
        if (p.find('#flag').is(':visible')) {
            flag++;
        } else {
            point.push(p);
        }
    }

    // 右
    p = $('#showTable tr').eq(trIndex).find('td').eq(tdIndex + 1);
    if (p.find('#flag').is(':visible')) {
        flag++;
    } else {
        point.push(p);
    }
    // 右下
    p = $('#showTable tr').eq(trIndex + 1).find('td').eq(tdIndex + 1);
    if (p.find('#flag').is(':visible')) {
        flag++;
    } else {
        point.push(p);
    }
    // 下
    p = $('#showTable tr').eq(trIndex + 1).find('td').eq(tdIndex);
    if (p.find('#flag').is(':visible')) {
        flag++;
    } else {
        point.push(p);
    }

    return {
        flag,
        point
    };
}

/** 以下為畫面效果 **/
// 滑鼠 hover
function onMap() {
    if ($(this).find('#content').is(':visible')) {
        return;
    }
    $(this).addClass('onMap');
}
// 滑鼠 hover over
function outMap() {
    if ($(this).find('#content').is(':visible')) {
        return;
    }
    $(this).removeClass();
    $(this).addClass('outMap');
}
// 遊戲結束提示
function isGameOver() {
    $("#myModal .modal-title").text('Game Over');
    $("#myModal").modal('show');
}
// 遊戲過關提示
function gameClear() {
    $("#myModal .modal-title").text('過關');
    $("#myModal").modal('show');
}
// 已開啟格子顯示
function openedShow(clickPoint) {
    var pNumber = clickPoint.find('#content').text().trim();
    var trIndex = clickPoint.closest('tr').index();
    var tdIndex = clickPoint.closest('td').index();
    var check = checkFlag(trIndex, tdIndex);

    if (pNumber != check.flag) {
        clickPoint.addClass('checkArount');

        setTimeout(function () {
            clickPoint.removeClass();
        }, 250);
    }
}
// 更改格子樣式
function print(color, point) {
    // 1-M 2-Number 3-Zero 4-Boom
    if (color == 1) {
        color = {
            'background': '#ff0000',
            'background': '-webkit-radial-gradient(left top, #ff0000, #b30000)',
            'background': '-o-linear-gradient(bottom right, #ff0000, #b30000)',
            'background': '-moz-linear-gradient(bottom right, #ff0000, #b30000)',
            'background': 'linear-gradient(to bottom right, #ff0000, #b30000)',
        };
        point.find('#flag').hide();
        point.find('#imgM').show();
        point.find('#content').text('');
    }
    if (color == 2) {
        color = {
            'background': '#0000cc',
            'background': '-webkit-radial-gradient(left top, #0000cc, #4d4dff)',
            'background': '-o-linear-gradient(bottom right, #0000cc, #4d4dff)',
            'background': '-moz-linear-gradient(bottom right, #0000cc, #4d4dff)',
            'background': 'linear-gradient(to bottom right, #0000cc, #4d4dff)',
        };
        point.find('#flag').hide();
    }
    if (color == 3) {
        color = {
            'background': '#99ccff',
            'background': '-webkit-radial-gradient(left top, #99ccff, #cce6ff)',
            'background': '-o-linear-gradient(bottom right, #99ccff, #cce6ff)',
            'background': '-moz-linear-gradient(bottom right, #99ccff, #cce6ff)',
            'background': 'linear-gradient(to bottom right, #99ccff, #cce6ff)',
        };
        point.find('#flag').hide();
        point.find('#content').text('');
    }
    if (color == 4) {
        color = {
            'background': '#99ccff',
            'background': '-webkit-radial-gradient(left top, #ffff00, #ffcc00)',
            'background': '-o-linear-gradient(bottom right, #ffff00, #ffcc00)',
            'background': '-moz-linear-gradient(bottom right, #ffff00, #ffcc00)',
            'background': 'linear-gradient(to bottom right, #ffff00, #ffcc00)',
        };
        point.find('#flag').hide();
        point.find('#imgM').hide();
        point.find('#content').prop('style', 'color: black');
        point.find('#icon').addClass('glyphicon glyphicon-fire');
        point.find('#icon').show();
    }

    if (color != 0) {
        point.css(color);
        if (point.find('#icon').is(':hidden')) {
            point.find('#content').show();
        }
    }
}