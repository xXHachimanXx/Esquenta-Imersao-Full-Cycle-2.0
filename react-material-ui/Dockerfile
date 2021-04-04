FROM node:14.15.4-alpine3.12

USER node

RUN mkdir -p /home/node/app && chown -R node:node /home/node/app
#RUN chmod -R /home/node/app
WORKDIR /home/node/app/

# Copiar tudo DESTE DIRETORIO para a IMAGEM
COPY . .

ENTRYPOINT [ "sh", ".docker/entrypoint.sh" ]
