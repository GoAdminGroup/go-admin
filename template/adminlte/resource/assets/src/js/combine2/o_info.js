$('.grid-per-pager').on('change', function (e) {
    console.log('changing...');
    $.pjax({url: this.value, container: '#pjax-container'});
});
$('.grid-refresh').on('click', function () {
    $.pjax.reload('#pjax-container');
    toastr.success('Refresh succeeded !');
});
toastr.options = {
    closeButton: true,
    progressBar: true,
    showMethod: 'slideDown',
    timeOut: 4000
};

$.pjax.defaults.timeout = 5000;
$.pjax.defaults.maxCacheLength = 0;
$(document).pjax('a:not(a[target="_blank"])', {
    container: '#pjax-container'
});

NProgress.configure({parent: '#pjax-container'});

$(document).on('pjax:timeout', function (event) {
    event.preventDefault();
});

$(document).on('submit', 'form[pjax-container]', function (event) {
    $.pjax.submit(event, '#pjax-container');
});

$(document).on('pjax:popstate', function () {

    $(document).one('pjax:end', function (event) {
        $(event.target).find('script[data-exec-on-popstate]').each(function () {
            $.globalEval(this.text || this.textContent || this.innerHTML || '');
        });
    });
});

$(document).on('pjax:send', function (xhr) {
    if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
        let submitBtn = $('form[pjax-container] :submit');
        if (submitBtn) {
            submitBtn.button('loading');
        }
    }
    NProgress.start();
});

$(document).on('pjax:complete', function (xhr) {
    if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
        let submitBtn = $('form[pjax-container] :submit');
        if (submitBtn) {
            submitBtn.button('reset');
        }
    }
    NProgress.done();
});

$(function () {
    $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
        let parent = $(this).parent().addClass('active');
        parent.siblings('.treeview.active').find('> a').trigger('click');
        parent.siblings().removeClass('active').find('li').removeClass('active');
    });

    $('[data-toggle="popover"]').popover();
});

let fullpageBtn = $('.fullpage-btn');
let exitFullpageBtn = $('.exit-fullpage-btn');

fullpageBtn.on('click', function () {
    launchFullscreen(document.documentElement);
    fullpageBtn.hide();
    exitFullpageBtn.show();
});

exitFullpageBtn.on('click', function () {
    exitFullscreen();
    exitFullpageBtn.hide();
    fullpageBtn.show();
});

function launchFullscreen(element) {
    if (element.requestFullscreen) {
        element.requestFullscreen();
    } else if (element.mozRequestFullScreen) {
        element.mozRequestFullScreen();
    } else if (element.msRequestFullscreen) {
        element.msRequestFullscreen();
    } else if (element.webkitRequestFullscreen) {
        element.webkitRequestFullScreen();
    }
}

function exitFullscreen() {
    if (document.exitFullscreen) {
        document.exitFullscreen();
    } else if (document.msExitFullscreen) {
        document.msExitFullscreen();
    } else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen();
    } else if (document.webkitExitFullscreen) {
        document.webkitExitFullscreen();
    }
}

$('.container-refresh').on('click', function () {
    $.pjax.reload('#pjax-container');
    toastr.success('Refresh succeeded !');
});

$('.sidebar-menu a').on('click', function () {

    let link = $(this).attr('href');
    if (link !== '#' && link.indexOf('http') === -1 && !checkNavExist(link)) {

        if (!checkNavLength()) {
            removeFirst();
        }

        removeActive();

        let addElement = $('<li class="active">\n' +
            '<a href="' + link + '">\n' +
            '<span>' + $(this).html().replace('<i/><span>', '<i/>&nbsp&nbsp&nbsp<span>') + '</span>\n' +
            '</a><i class="close-tab fa fa-remove"></i>\n' +
            '</li>');

        addElement.find('.close-tab').on('click', function () {
            let li = $(this).parent();
            if (li.hasClass('active')) {
                if (li.prev().length > 0) {
                    li.prev().addClass('active');
                    $.pjax({url: li.prev().find('a').attr('href'), container: '#pjax-container'});
                } else if (li.next().length > 0) {
                    li.next().addClass('active');
                    $.pjax({url: li.next().find('a').attr('href'), container: '#pjax-container'});
                }
            }
            li.remove();
        });
        addElement.on('mouseover', function () {
            if ($(this).children('i')) {
                $(this).children('i').show();
            }
        });
        addElement.on('mouseout', function () {
            if ($(this).children('i')) {
                $(this).children('i').hide();
            }
        });
        addElement.on('click', function () {
            removeActive();
            $(this).addClass('active');
        });

        addElement.appendTo('.nav-addtabs');
    }
});

function checkNavExist(link) {
    let navs = $('.nav-addtabs li');
    for (let i = 0; i < navs.length; i++) {
        if ($(navs[i]).find('a').attr('href') === link) {
            removeActive();
            $(navs[i]).addClass('active');
            return true;
        }
    }
    return false;
}

function removeFirst() {
    let navs = $('.nav-addtabs li');
    $(navs[0]).remove();
}

function removeActive() {
    let lis = $('.nav-addtabs li');
    for (let i = 0; i < lis.length; i++) {
        $(lis[i]).removeClass('active');
    }
}

function checkNavLength() {
    return $('#firstnav').width() * 0.6 - $('.nav-addtabs').width() >= 120;
}

$('.fixed-btn').on('click', function () {
    let clicked = $(this).attr('data-click');
    if (clicked === "false") {
        $('.main-sidebar').css('position', 'fixed');
        $('.main-header .logo').css('position', 'fixed');
        $(this).attr('data-click', 'true')
    } else {
        $('.main-sidebar').css('position', '');
        $('.main-header .logo').css('position', '');
        $(this).attr('data-click', 'false')
    }
});
