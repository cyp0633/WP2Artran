# WP2Artran - 转换 WordPress 评论为 Artran

[English](./README.md)

与传统的导出方法相比，REST API 显然是一种更现代的获取 WordPress 评论的方式。另一方面，Artrans 是 [Artalk](https://artalk.js.org/) 使用的评论格式，可以导入到其他一些评论系统中，例如 Twikoo。这使得 Artrans 成为一种很好的中间评论格式。

我们假设你的新旧站点都已经上线，并且你已经设置了新站点的重定向。例如，旧的 WordPress 站点上有一篇文章位于 <https://old.example.com/archives/123>，对应的新文章位于 <https://new.example.com/post/new-post>。"设置重定向" 意味着当有人访问 <https://new.example.com/archives/123> 时，浏览器会被重定向到 <https://new.example.com/post/new-post>。你不需要在新站点上提前设置 Artalk 或任何其他东西。

你需要一个配置文件来使用这个工具。我们继续使用上面的例子。在这种情况下，你的配置如下：

```yaml
old:
  hostname: "old.example.com"
  permalink-prefix: "/archives/"

new:
    hostname: "new.example.com"
    permalink-prefix: "/post/"

output-path: "./comments.json.artrans"

auth:
  username:
  password:
```

你**必须**通过 `output-path` 字段指定输出 Artrans 文件路径。

`auth` 字段的内容来自 WordPress 应用密码。要生成密码，请参考[官方文档中的 "Generating Manually" 部分](https://make.wordpress.org/core/2020/11/05/application-passwords-integration-guide/#generating-manually)。或者你可以使用任何 REST API 支持的凭证对。

设置好这些后，运行以下命令：

```bash
$ ./WP2Artran config.yaml
```

程序将自动通过 WordPress REST API 获取评论，解析重定向，并将其转换为 Artrans。如果重定向解析失败，程序将打印一个警告并继续转换评论。在这种情况下，该 Artran 的 `PageKey` 字段是旧的文章 ID（在例子中是 `/123/`）。这种情况经常发生在导出 WordPress 的“页面”类型时。
