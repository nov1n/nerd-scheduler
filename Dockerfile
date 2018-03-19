FROM scratch
ADD nerd-scheduler /nerd-scheduler
ENTRYPOINT ["/nerd-scheduler"]
