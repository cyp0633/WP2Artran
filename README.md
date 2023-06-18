# WP2Artran - WordPress to Artalk comment format

[简体中文 zh-CN](./README.zh-CN.md)

Compared to the legacy export method, REST API is obviously a more modern way to retrieve WordPress comments. On the other hand, Artrans is a comment format used by [Artalk](https://artalk.js.org/) and can be imported into some other comment systems, e.g. Twikoo. This makes Artrans a good intermediate comment format.

We assume that both your old and new site is online, and you've set up redirections of the new site. For example, there's an old post on a WordPress instance at <https://old.example.com/archives/123>, and the corresponding new post is at <https://new.example.com/post/new-post>. "Set up redirections" means that when someone visits <https://new.example.com/archives/123>, the browser will be redirected to <https://new.example.com/post/new-post>. You don't have to set Artalk or anything else up on the new site.

You need a configuration file to use this tool. We continue to take the example above. In this case, your configuration is as below:

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

You MUST specify output Artrans file path by `output-path` field. 

The content of `auth` field involves in WordPress application password. To generate the password, please refer to "Generating Manually" section in [the official documentation](https://make.wordpress.org/core/2020/11/05/application-passwords-integration-guide/#generating-manually). Or you could use any credential pair that's supported by REST API authentication method.

After setting all those up, run the following command:

```bash
$ ./WP2Artran config.yaml
```

The app will automatically fetch comments via WordPress REST API, resolve redirections, and convert it to Artrans. If redirection resolution fails, the app will print a warning and continue to convert the comment. In that case, the `PageKey` field of the Artran is the old post ID (`/123/` in the example).
