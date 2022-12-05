package quic

//go:generate sh -c "./mockgen_private.sh quic mock_send_conn_test.go github.com/kixelated/quic-go sendConn"
//go:generate sh -c "./mockgen_private.sh quic mock_sender_test.go github.com/kixelated/quic-go sender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_internal_test.go github.com/kixelated/quic-go streamI"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_stream_test.go github.com/kixelated/quic-go cryptoStream"
//go:generate sh -c "./mockgen_private.sh quic mock_receive_stream_internal_test.go github.com/kixelated/quic-go receiveStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_send_stream_internal_test.go github.com/kixelated/quic-go sendStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_sender_test.go github.com/kixelated/quic-go streamSender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_getter_test.go github.com/kixelated/quic-go streamGetter"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_data_handler_test.go github.com/kixelated/quic-go cryptoDataHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_frame_source_test.go github.com/kixelated/quic-go frameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_ack_frame_source_test.go github.com/kixelated/quic-go ackFrameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_manager_test.go github.com/kixelated/quic-go streamManager"
//go:generate sh -c "./mockgen_private.sh quic mock_sealing_manager_test.go github.com/kixelated/quic-go sealingManager"
//go:generate sh -c "./mockgen_private.sh quic mock_unpacker_test.go github.com/kixelated/quic-go unpacker"
//go:generate sh -c "./mockgen_private.sh quic mock_packer_test.go github.com/kixelated/quic-go packer"
//go:generate sh -c "./mockgen_private.sh quic mock_mtu_discoverer_test.go github.com/kixelated/quic-go mtuDiscoverer"
//go:generate sh -c "./mockgen_private.sh quic mock_conn_runner_test.go github.com/kixelated/quic-go connRunner"
//go:generate sh -c "./mockgen_private.sh quic mock_quic_conn_test.go github.com/kixelated/quic-go quicConn"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_test.go github.com/kixelated/quic-go packetHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_unknown_packet_handler_test.go github.com/kixelated/quic-go unknownPacketHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_manager_test.go github.com/kixelated/quic-go packetHandlerManager"
//go:generate sh -c "./mockgen_private.sh quic mock_multiplexer_test.go github.com/kixelated/quic-go multiplexer"
//go:generate sh -c "./mockgen_private.sh quic mock_batch_conn_test.go github.com/kixelated/quic-go batchConn"
//go:generate sh -c "mockgen -package quic -self_package github.com/kixelated/quic-go -destination mock_token_store_test.go github.com/kixelated/quic-go TokenStore"
//go:generate sh -c "mockgen -package quic -self_package github.com/kixelated/quic-go -destination mock_packetconn_test.go net PacketConn"
