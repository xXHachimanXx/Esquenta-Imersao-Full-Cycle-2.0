<<<<<<< HEAD
FROM node:14.15.4-alpine3.12

USER node

RUN mkdir -p /home/node/app && chown -R node:node /home/node/app
WORKDIR /home/node/app/

# Copiar tudo DESTE DIRETORIO para a IMAGEM
COPY . .

ENTRYPOINT [ "sh", ".docker/entrypoint.sh" ]
=======
FROM node:14.15.4-alpine3.12

USER node

RUN mkdir -p /home/node/app && chown -R node:node /home/node/app
#RUN chmod -R /home/node/app
WORKDIR /home/node/app/

# Copiar tudo DESTE DIRETORIO para a IMAGEM
COPY . .

ENTRYPOINT [ "sh", ".docker/entrypoint.sh" ]
>>>>>>> 9730c3e8e238fcf70168252ad0cf984e24521105
