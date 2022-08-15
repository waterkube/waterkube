FROM gcr.io/distroless/base-debian11
ARG TARGETARCH

WORKDIR /app

COPY dist/.env ./
COPY dist/${TARGETARCH}/waterkube ./

EXPOSE 4000

ENTRYPOINT ["./waterkube"]
CMD ["web", "serve"]
