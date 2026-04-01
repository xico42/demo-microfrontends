FROM php:8-apache
RUN echo '<Directory /var/www/html>\n    FallbackResource /index.php\n</Directory>' \
    > /etc/apache2/conf-enabled/fallback.conf
COPY index.php /var/www/html/index.php
