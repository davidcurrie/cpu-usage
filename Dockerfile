FROM lalyos/scratch-chmx
ADD https://github.com/davidcurrie/cpu-usage/releases/download/v1.0/cpu-usage /bin/cpu-usage
RUN ["/bin/chmx", "/bin/cpu-usage"]
CMD ["/bin/cpu-usage"]
