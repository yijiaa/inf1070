FROM moussaabdenbi/inf1070:tp2-h2021

# Mettre à jour les paquets de l'image
RUN apt-get update -y

# Installer elinks
RUN apt-get install -y elinks

# Copier le répertoire s
COPY ./s /go/src/github.com/zquestz/s

# Se placer dans le bon répertoire
WORKDIR /go/src/github.com/zquestz/s

# Compiler l'application
RUN make

# Installer l'application
RUN make install

# Créer le répertoire /root/.config/s
RUN mkdir -p /root/.config/s

