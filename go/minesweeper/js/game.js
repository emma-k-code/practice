$(document).ready(init);

var clickCount = 0;

var ws = new WebSocket("ws://127.0.0.1:8080/ws");

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
ws.onmessage = function (event) {
    var data = event.data;
    $('#showTable').html(data);
    $('#showTable td').click(clickMap);
    $('#showTable td').dblclick(dbClickMap);
    $('#showTable td').hover(onMap, outMap);
    $('#showTable td').mousedown(function (event) {
        if ($(this).find('#content').is(':hidden')) {
            switch (event.which) {
                case 3:
                    $(this).find('#flag').toggle();
                    break;
            }
        }
    });
};

function setTable() {
    param = $("#row").val() + "," + $("#column").val() + "," + $("#m").val()
    ws.send(param);
}

function resetTable(trIndex, tdIndex) {
    $('#showTable tr').eq(trIndex).find('td').eq(tdIndex).click();
    param = $("#row").val() + "," + $("#column").val() + "," + $("#m").val()
    ws.send(param);
}

function onMap() {
    if ($(this).find('#content').is(':visible')) {
        return;
    }
    $(this).addClass('onMap');
}

function outMap() {
    if ($(this).find('#content').is(':visible')) {
        return;
    }
    $(this).removeClass();
    $(this).addClass('outMap');
}

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

function dbClickMap() {
    if ($(this).find('#content').is(':hidden')) {
        return;
    }

    var pNumber = $(this).find('#content').text().trim();
    var trIndex = $(this).closest('tr').index();
    var tdIndex = $(this).closest('td').index();
    var check = checkFlag(trIndex, tdIndex);
    if (pNumber == check.flag) {
        aroundPoint(trIndex, tdIndex);
    } else {
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
    }

}

function clickMap() {

    if ($(this).find('#flag').is(':visible')) {
        return;
    }

    if ($(this).find('#content').is(':visible')) {
        var pNumber = $(this).find('#content').text().trim();
        var trIndex = $(this).closest('tr').index();
        var tdIndex = $(this).closest('td').index();
        var check = checkFlag(trIndex, tdIndex);
        var point = $(this);

        if (pNumber != check.flag) {
            point.addClass('checkArount');

            setTimeout(function () {
                point.removeClass();
            }, 250);
        }
        return;
    }

    var gameover = true;
    $('#showTable td').each(function () {
        if ($(this).find('#content').is(':hidden')) {
            gameover = false;
        }
    });

    if (gameover) {
        return;
    }

    if ($(this).find('#content').text().trim() == "M") {
        if (clickCount == 0) {
            var trIndex = $(this).closest('tr').index();
            var tdIndex = $(this).closest('td').index();
            resetTable(trIndex, tdIndex);
            clickCount++;
            return;
        } else {
            checkOver($(this));
            gameover = true;
        }
    }

    if (gameover) {
        return;
    }

    $(this).removeClass();

    print(2, $(this));

    if ($(this).find('#content').text().trim() == '') {
        var trIndex = $(this).closest('tr').index();
        var tdIndex = $(this).closest('td').index();

        print(3, $(this));

        aroundPoint(trIndex, tdIndex);
    }

    if (checkPass()) {
        $('#showTable td').each(function () {
            $(this).find('#flag').hide();
            var color = checkPoint($(this));
            print(color, $(this));
        });

        $("#myModal .modal-title").text('過關');
        $("#myModal").modal('show');
    }

    clickCount++;
}

function checkPass() {
    gameover = true;
    $('#showTable td').each(function () {
        if ($(this).find('#content').is(':hidden')) {
            gameover = false;
        }
    });
    if (gameover) {
        return;
    }
    var i = 0
    $('#showTable td').each(function () {
        if ($(this).find('#content').is(':hidden')) {
            if ($(this).find('#content').text().trim() != "M") {
                i++;
            }
        }
    });

    if (i == 0) {
        return true;
    }
}

function checkOver(point) {
    $('#showTable td').each(function () {
        if ($(this).find('#flag').is(':visible')) {
            $(this).find('#flag').hide();
            if ($(this).find('#content').text().trim() != "M") {
                $(this).find('#content').text('');
                $(this).find('#content').show();
                $(this).find('#icon').addClass('glyphicon glyphicon-remove');
                $(this).find('#icon').show();
            }
        }
        var color = checkPoint($(this));
        print(color, $(this));
    });

    print(4, point);
    $("#myModal .modal-title").text('Game Over');
    $("#myModal").modal('show');
}

function aroundPoint(trIndex, tdIndex) {
    var point = [];
    var zeroPoint = [];

    // 判斷是否為最上方
    if (trIndex != 0) {
        // 上
        p = $('#showTable tr').eq(trIndex - 1).find('td').eq(tdIndex);
        if (p.find('#content').is(':hidden')) {
            if (p.find('#content').text().trim() == '') {
                zeroPoint.push(p);
            } else {
                point.push(p);
            }
        }


        // 判斷是否為最左側
        if (tdIndex != 0) {
            // 左上
            p = $('#showTable tr').eq(trIndex - 1).find('td').eq(tdIndex - 1);
            if (p.find('#content').is(':hidden')) {
                if (p.find('#content').text().trim() == '') {
                    zeroPoint.push(p);
                } else {
                    point.push(p);
                }
            }
        }

        // 右上
        p = $('#showTable tr').eq(trIndex - 1).find('td').eq(tdIndex + 1);
        if (p.find('#content').is(':hidden')) {
            if (p.find('#content').text().trim() == '') {
                zeroPoint.push(p);
            } else {
                point.push(p);
            }
        }
    }

    // 判斷是否為最左側
    if (tdIndex != 0) {
        // 左
        p = $('#showTable tr').eq(trIndex).find('td').eq(tdIndex - 1);
        if (p.find('#content').is(':hidden')) {
            if (p.find('#content').text().trim() == '') {
                zeroPoint.push(p);
            } else {
                point.push(p);
            }
        }
        // 左下
        p = $('#showTable tr').eq(trIndex + 1).find('td').eq(tdIndex - 1);
        if (p.find('#content').is(':hidden')) {
            if (p.find('#content').text().trim() == '') {
                zeroPoint.push(p);
            } else {
                point.push(p);
            }
        }
    }

    // 右
    p = $('#showTable tr').eq(trIndex).find('td').eq(tdIndex + 1);
    if (p.find('#content').is(':hidden')) {
        if (p.find('#content').text().trim() == '') {
            zeroPoint.push(p);
        } else {
            point.push(p);
        }
    }
    // 右下
    p = $('#showTable tr').eq(trIndex + 1).find('td').eq(tdIndex + 1);
    if (p.find('#content').is(':hidden')) {
        if (p.find('#content').text().trim() == '') {
            zeroPoint.push(p);
        } else {
            point.push(p);
        }
    }
    // 下
    p = $('#showTable tr').eq(trIndex + 1).find('td').eq(tdIndex);
    if (p.find('#content').is(':hidden')) {
        if (p.find('#content').text().trim() == '') {
            zeroPoint.push(p);
        } else {
            point.push(p);
        }
    }


    openAround(point);
    openZeroAround(zeroPoint);

    if (checkPass()) {
        $('#showTable td').each(function () {
            $(this).find('#flag').hide();
            var color = checkPoint($(this));
            print(color, $(this));
        });
        $("#myModal .modal-title").text('過關');
        $("#myModal").modal('show');
    }
}

function openAround(point) {
    $.each(point, function () {
        if ($(this).find('#flag').is(':hidden')) {
            if ($(this).find('#content').is(':hidden')) {
                if ($(this).find('#content').text().trim() == "M") {
                    checkOver($(this));
                } else {
                    print(2, $(this));
                }
            }
        }
    });
}

function openZeroAround(point) {
    $.each(point, function () {
        if ($(this).find('#flag').is(':hidden')) {
            if ($(this).find('#content').is(':hidden')) {
                print(3, $(this));
                aroundPoint($(this).closest('tr').index(), $(this).closest('td').index());
            }
        }
    });
}

function checkPoint(point) {
    if (point.find('#flag').is(':hidden')) {
        if (point.find('#content').text().trim() == '') {
            return 3;
        } else if (point.find('#content').text().trim() == "M") {
            return 1;
        } else {
            return 2;
        }
    }

    return 0;
}

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
    }
    if (color == 3) {
        color = {
            'background': '#99ccff',
            'background': '-webkit-radial-gradient(left top, #99ccff, #cce6ff)',
            'background': '-o-linear-gradient(bottom right, #99ccff, #cce6ff)',
            'background': '-moz-linear-gradient(bottom right, #99ccff, #cce6ff)',
            'background': 'linear-gradient(to bottom right, #99ccff, #cce6ff)',
        };
    }
    if (color == 4) {
        color = {
            'background': '#99ccff',
            'background': '-webkit-radial-gradient(left top, #ffff00, #ffcc00)',
            'background': '-o-linear-gradient(bottom right, #ffff00, #ffcc00)',
            'background': '-moz-linear-gradient(bottom right, #ffff00, #ffcc00)',
            'background': 'linear-gradient(to bottom right, #ffff00, #ffcc00)',
        };
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