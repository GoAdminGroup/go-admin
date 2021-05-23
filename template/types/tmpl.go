package types

var tmpls = map[string]string{"choose": `{{define "choose"}}
    <script>
        $("select.{{.Field}}").on("select2:select", function (e) {
            if (e.params.data.text === {{.Val}} || e.params.data.id === {{.Val}}) {
                if ($("select.{{.ChooseField}}").length > 0) {
                    $("select.{{.ChooseField}}").val("{{.Value}}[0]").select2()
                } else {
                    $(".{{.ChooseField}}").val({{.Value}})
                }
            }
        })
    </script>
{{end}}`, "choose_ajax": `{{define "choose_ajax"}}
    <script>

        let {{.Field}}_updateBoxSelections = function (selectObj, new_opts) {
            selectObj.html('');
            new_opts.forEach(function (opt) {
                selectObj.append($('<option value="' + opt["id"] + '">' + opt["text"] + '</option>'));
            });
            selectObj.bootstrapDualListbox('refresh', true);
        };

        let {{.Field}}_req = function (selectObj, box, event) {
            $.ajax({
                url: "{{.Url}}",
                type: 'post',
                dataType: 'text',
                data: {
                    'value': $("select.{{.Field}}").val(),
                    {{.PassValue}}
                    'event': event
                },
                success: function (data) {
                    if (typeof (data) === "string") {
                        data = JSON.parse(data);
                    }
                    if (data.code === 0) {
                        {{if eq .ActionJS ""}}

                        if (selectObj.length > 0) {
                            if (typeof (data.data) === "object") {
                                if (box) {
                                    {{.Field}}_updateBoxSelections(selectObj, data.data)
                                } else {
                                    if (typeof (selectObj.attr("multiple")) !== "undefined") {
                                        selectObj.html("");
                                    }
                                    selectObj.select2({
                                        data: data.data
                                    });
                                }
                            } else {
                                if (box) {
                                    selectObj.val(data.data).select2()
                                } else {

                                }
                            }
                        } else {
                            $('.{{.ChooseField}}').val(data.data);
                        }

                        {{else}}

                        {{.ActionJS}}

                        {{end}}
                    } else {
                        swal(data.msg, '', 'error');
                    }
                },
                error: function () {
                    alert('error')
                }
            });
        };

        if ($("label[for='{{.Field}}']").next().find(".bootstrap-duallistbox-container").length === 0) {
            $("select.{{.Field}}").on("select2:select", function (e) {
                let id = '{{.ChooseField}}';
                let selectObj = $("select." + id);
                if (selectObj.length > 0) {
                    selectObj.val("").select2();
                    selectObj.html('<option value="" selected="selected"></option>')
                }
                {{.Field}}_req(selectObj, false, "select");
            });
            if (typeof ($("select.{{.Field}}").attr("multiple")) !== "undefined") {
                $("select.{{.Field}}").on("select2:unselect", function (e) {
                    let id = '{{.ChooseField}}';
                    let selectObj = $("select." + id);
                    if (selectObj.length > 0) {
                        selectObj.val("").select2();
                        selectObj.html('<option value="" selected="selected"></option>')
                    }
                    {{.Field}}_req(selectObj, false, "unselect");
                })
            }
        } else {
            let {{.Field}}_lastState = $(".{{.Field}}").val();

            $(".{{.Field}}").on('change', function (e) {
                var newState = $(this).val();
                if ($({{.Field}}_lastState).not(newState).get().length > 0) {
                    let id = '{{.ChooseField}}';
                    {{.Field}}_req($("." + id), true, "unselect");
                }
                if ($(newState).not({{.Field}}_lastState).get().length > 0) {
                    let id = '{{.ChooseField}}';
                    {{.Field}}_req($("." + id), true, "select");
                }
                {{.Field}}_lastState = newState;
            })
        }
    </script>
{{end}}`, "choose_custom": `{{define "choose_custom"}}
    <script>
        $("select.{{.Field}}").on("select2:select", function (e) {
            {{.JS}}
        })
    </script>
{{end}}`, "choose_disable": `{{define "choose_disable"}}
    <script>
        $("select.{{.Field}}").on("select2:select", function (e) {
            if ({{.Value}}.indexOf(e.params.data.text) !== -1 || {{.Value}}.indexOf(e.params.data.id) !== -1) {
                {{range $key, $fields := .ChooseFields}}

                $(".{{$fields}}").prop('disabled', true);

                {{end}}
            } else {
                {{range $key, $fields := .ChooseFields}}

                $(".{{$fields}}").prop('disabled', false);

                {{end}}
            }
        });
    </script>
{{end}}`, "choose_hide": `{{define "choose_hide"}}
    <script>
        $("select.{{.Field}}").on("select2:select", function (e) {
            if ({{.Value}}.indexOf(e.params.data.text) !== -1 || {{.Value}}.indexOf(e.params.data.id) !== -1) {
                {{range $key, $fields := .ChooseFields}}

                $("label[for='{{$fields}}']").parent().hide();

                {{end}}
            } else {
                {{range $key, $fields := .ChooseFields}}

                $("label[for='{{$fields}}']").parent().show();

                {{end}}
            }
        });
        $(function () {
            let {{.Field}}data = $(".{{.Field}}").select2("data");
            let {{.Field}}text = "";
            let {{.Field}}id = "";
            if ({{.Field}}data.length > 0) {
                {{.Field}}text = {{.Field}}data[0].text;
                {{.Field}}id = {{.Field}}data[0].id;
            }
            if ({{.Value}}.indexOf({{$.Field}}text) !== -1 || {{.Value}}.indexOf({{$.Field}}id) !== -1) {
                {{range $key, $fields := .ChooseFields}}

                $("label[for='{{$fields}}']").parent().hide();

                {{end}}
            }
        })
    </script>
{{end}}`, "choose_map": `{{define "choose_map"}}
    <script>
        $("select.{{.Field}}").on("select2:select", function (e) {
            {{range $val, $object := .Data}}

            {{if $object.Hide}}

            if (e.params.data.text === "{{$val}}" || e.params.data.id === "{{$val}}") {
                $("label[for='{{$object.Field}}']").parent().hide()
            } else {
                $("label[for='{{$object.Field}}']").parent().show()
            }

            {{else if $object.Disable}}

            if (e.params.data.text === "{{$val}}" || e.params.data.id === "{{$val}}") {
                $("label[for='{{$object.Field}}']").prop('disabled', true);
            } else {
                $("label[for='{{$object.Field}}']").prop('disabled', false);
            }

            {{else}}

            if (e.params.data.text === "{{$val}}" || e.params.data.id === "{{$val}}") {
                if ($("select.{{$object.Field}}").length > 0) {
                    $("select.{{$object.Field}}").val("{{$object.Value}}").select2()
                } else {
                    $("#{{$object.Field}}").val("{{$object.Value}}")
                }
            }

            {{end}}

            {{end}}
        })
    </script>
{{end}}`, "choose_show": `{{define "choose_show"}}
    <script>
        $("select.{{.Field}}").on("select2:select", function (e) {
            if ({{.Value}}.indexOf(e.params.data.text) !== -1 || {{.Value}}.indexOf(e.params.data.id) !== -1) {
                {{range $key, $fields := .ChooseFields}}

                $("label[for='{{$fields}}']").parent().show();

                {{end}}
            } else {
                {{range $key, $fields := .ChooseFields}}

                $("label[for='{{$fields}}']").parent().hide();

                {{end}}
            }
        });
        $(function () {
            let {{.Field}}data = $(".{{.Field}}").select2("data");
            let {{.Field}}text = "";
            let {{.Field}}id = "";
            if ({{.Field}}data.length > 0) {
                {{.Field}}text = {{.Field}}data[0].text;
                {{.Field}}id = {{.Field}}data[0].id;
            }
            if ({{.Value}}.indexOf({{$.Field}}text) !== -1 || {{.Value}}.indexOf({{$.Field}}id) !== -1) {
                {{range $key, $fields := .ChooseFields}}

                $("label[for='{{$fields}}']").parent().show();

                {{end}}
            }
        })
    </script>
{{end}}`}
