#
#compiling stage
#
FROM golang AS compiling_stage
RUN mkdir -p /go/src/pipeline
WORKDIR /go/src/pipeline
#
#add all sources
ADD bufferizationStage.go .
ADD consoleDataSource.go .
ADD cycleIntBuffer.go .
ADD intBufferNode.go .
ADD intConsoleReader.go .
ADD interfaces.go .
ADD intRingBuffer.go .
ADD main.go .
ADD modFiltration.go .
ADD negativeFiltrationStage.go .
ADD pipeline.go .
#add module
ADD go.mod .
#build and install
RUN go install .
#
#stage2: lightwieght image copy 
#
FROM alpine:latest
LABEL version="1.1.0"
WORKDIR /root/
COPY --from=compiling_stage /go/bin/pipelinemx .
ENTRYPOINT ./pipelinemx