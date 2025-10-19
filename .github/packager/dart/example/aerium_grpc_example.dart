import 'package:aerium_grpc/aerium_grpc.dart';
import 'package:grpc/grpc.dart';

Future<void> main(List<String> arguments) async {
  final channel = ClientChannel(
    'localhost',
    port: 50051,
    options: ChannelOptions(
      credentials: ChannelCredentials.insecure(),
    ),
  );

  final networkClient = NetworkServiceClient(channel);

  try {
    final response = await networkClient.getNetworkInfo(GetNetworkInfoRequest());
    print('Network name: ${response.networkName}');
  } catch (e) {
    print('Caught error: $e');
  }

  await channel.shutdown();
}
