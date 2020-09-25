#! /bin/sh

#
# Crea el usuario y la base de datos (leer README.txt para instrucciones de instalacion)
#

# Crea el usuario para la base de datos
sudo -u postgres /usr/lib/postgresql/11/bin/createuser -d ironuser

# Crea la base de datos
sudo -u postgres /usr/lib/postgresql/11/bin/createdb -O ironuser -E=UTF-8 --lc-collate=en_US.UTF-8 --lc-ctype=en_US.UTF-8 irondb