package main

func main() {
	/*
		// Загрузка клиентских сертификатов и CA сертификата
		clientCert, err := tls.LoadX509KeyPair("./cert/client/client.crt", "./cert/client/client.key")
		if err != nil {
			log.Fatalf("failed to load client certificate: %v", err)
		}

		file, err := os.Open("./cert/server/ca.crt")
		if err != nil {
			log.Fatalf("failed to open CA certificate file: %v", err)
		}
		defer file.Close()

		caCert, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("failed to read CA certificate: %v", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      caCertPool,
		})

		// Настройка gRPC клиента с TLS
		conn, err := grpc.Dial(":3300", grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		slog.Info("connected to server")
	*/
}
