# Usar uma imagem base com Node.js
FROM node:18-alpine

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o package.json e o package-lock.json para o container
COPY package.json package-lock.json ./

# Instalar dependências
RUN npm install --legacy-peer-deps

# Copiar todo o código para o container
COPY . .

# Construir o projeto Next.js
RUN npm run build

# Expor a porta usada pelo Next.js
EXPOSE 3000

# Comando para iniciar o servidor
CMD ["npm", "start"]
