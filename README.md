# Webhook-Forward

![Golang](https://img.shields.io/github/workflow/status/starudream/webhook-forward/Golang/master?style=for-the-badge)
![Docker](https://img.shields.io/github/workflow/status/starudream/webhook-forward/Docker/master?style=for-the-badge)
![License](https://img.shields.io/badge/License-Apache%20License%202.0-blue?style=for-the-badge)

## Usage

![Version](https://img.shields.io/docker/v/starudream/webhook-forward?style=for-the-badge)
![Size](https://img.shields.io/docker/image-size/starudream/webhook-forward/latest?style=for-the-badge)
![Pull](https://img.shields.io/docker/pulls/starudream/webhook-forward?style=for-the-badge)

```bash
docker pull starudream/webhook-forward
```

```bash
docker run -d starudream/webhook-forward
```

### Env

```shell
ADDR=127.0.0.1:9988
PROXY=http://127.0.0.1:7890
DEBUG=true

REDIS_DSN=redis://127.0.0.1:6378 #optional

DINGTALK_ENABLE=true
DINGTALK_TOKEN=
DINGTALK_SECRET=

TELEGRAM_ENABLE=true
TELEGRAM_TOKEN=
TELEGRAM_TO=

WEIXIN_ENABLE=true
WEIXIN_ID=
WEIXIN_SECRET=
WEIXIN_AID=
WEIXIN_TO=@all
```

## License

[Apache License 2.0](./LICENSE)
