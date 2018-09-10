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

$(document).on("pjax:popstate", function () {

    $(document).one("pjax:end", function (event) {
        $(event.target).find("script[data-exec-on-popstate]").each(function () {
            $.globalEval(this.text || this.textContent || this.innerHTML || '');
        });
    });
});

$(document).on('pjax:send', function (xhr) {
    if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
        $submit_btn = $('form[pjax-container] :submit');
        if ($submit_btn) {
            $submit_btn.button('loading')
        }
    }
    NProgress.start();
});

$(document).on('pjax:complete', function (xhr) {
    if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
        $submit_btn = $('form[pjax-container] :submit');
        if ($submit_btn) {
            $submit_btn.button('reset')
        }
    }
    NProgress.done();
});

$(function () {
    $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
        var $parent = $(this).parent().addClass('active');
        $parent.siblings('.treeview.active').find('> a').trigger('click');
        $parent.siblings().removeClass('active').find('li').removeClass('active');
    });

    $('[data-toggle="popover"]').popover();
});

function getQueryString(name) {
    let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    let r = window.location.search.substr(1).match(reg);
    if (r != null) {
        return decodeURI(r[2]);
    }
    return null;
}

$(function () {
    sortType = getQueryString('sort_type');
    if (sortType != null && $('a.fa.fa-fw').length > 0) {
        if (sortType === "asc") {
            $('a.fa.fa-fw').removeClass("fa-sort-amount-desc");
            $('a.fa.fa-fw').addClass("fa-sort-amount-asc");
            href = $('a.fa.fa-fw').attr("href");
            $('a.fa.fa-fw').attr("href", href.replace("asc", "desc"));
        } else {
            $('a.fa.fa-fw').removeClass("fa-sort-amount-asc");
            $('a.fa.fa-fw').addClass("fa-sort-amount-desc");
            href = $('a.fa.fa-fw').attr("href");
            $('a.fa.fa-fw').attr("href", href.replace("desc", "asc"));
        }
    }
});


if ($('.grid-select-all').length > 0) {
    $('.grid-row-delete').unbind('click').click(function () {
        DeletePost($(this).data('id'));
    });

    let selectedRows = function () {
        let selected = [];
        $('.grid-row-checkbox:checked').each(function(){
            selected.push($(this).data('id'));
        });
        return selected;
    };

    $('.grid-select-all').on('ifChanged', function (event) {
        if (this.checked) {
            $('.grid-row-checkbox').iCheck('check');
        } else {
            $('.grid-row-checkbox').iCheck('uncheck');
        }
    });
    $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'});

    $(function () {
        $('.grid-row-checkbox').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
            if (this.checked) {
                $(this).closest('tr').css('background-color', "#ffffd5");
            } else {
                $(this).closest('tr').css('background-color', '');
            }
        });
    });

    $('.grid-batch-0').on('click', function() {
        DeletePost(selectedRows().join());
    });
}