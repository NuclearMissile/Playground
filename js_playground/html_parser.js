var str = `<html>
<head class="aaa"><title>Hello</title></head>
<body>
<div id="container"><br />
<div class="header">
</div>
<div class="content">
</div>
<div class="footer">
</div>
</div>
</body>
</html>`;

 function parser(str) {
    var nodes = [];
    //第一步，排除所有属性值，防止里面出现<>引起干扰
    var attrCached = {},
        index = 1;
    str = str.replace(/"\w+"/g, function(str) {
        var key = "&" + index;
        index++;
        attrCached[key] = str.slice(1, -1);
        return key;
    });
    //处理开标签
    var tagCached = {},
        tagIndex = 1;
    str = str.replace(/<(\w+)([^>]+)>/g, function(str, tag, attrs) {
        var propsValue = {},
            selfClose = false;
        if (attrs[attrs.length - 1] == "/") {
            selfClose = true;

            attrs = attrs.slice(0, -1);
        }
        if (attrs[0] !== " ") {
            tag += attrs;
        } else {
            attrs.replace(/\S+/g, function(attr) {
                var arr = attr.split("=");
                propsValue[arr[0]] = attrCached[arr[1]];
            });
        }
        var key = "&" + tagIndex;
        var obj = (tagCached[key] = {
            tag: tag,
            props: propsValue,
            children: []
        });
        if (selfClose) {
            obj.selfClose = true;
        }
        tagIndex++;
        return key;
    });
    var numCache = {};
    for (var num = 0; num <= 9; num++) {
        numCache[num] = 1;
    }
    var parentStack = [];
    var root = {};
    var parent = null;
    for (var i = 0; i < str.length; i++) {
        if (str[i] == "&") {
            var key = "&";
            var j = i + 1;
            while (numCache[str[j]]) {
                key += str[j];
                j++;
            }
            i = j - 1;
            var tagObj = tagCached[key];
            tagObj.parent = parent;
            parentStack.push(tagObj);
            if (!tagObj.selfClose) {
                parent = tagObj;
                parentStack = tagObj.children;
            }
        } else if (str[i] == "<") {
            if (str[i + 1] == "\/") {
                var j = i + 2;
                var endTag = "";
                while (str[j] != ">") {
                    endTag += str[j];
                    j++;
                }
                i = j;
                if (parent && parent.tag == endTag) {
                    parent = parent.parent;
                    parentStack = parent && parent.children;
                }
            } else {
                throw "没有闭合标签";
            }
        } else if (str[i] != ">") {
            var text = "";
            var j = i;
            while (str[j] !== "&" && str[j] !== "<" && str[j]) {
                text += str[j];
                j++;
            }
            i = j - 1;
            if (text.trim().length !== 0) {
                parentStack.push(text);
            }
        }
    }

    return tagCached["&1"];
}
console.log(parser(str));