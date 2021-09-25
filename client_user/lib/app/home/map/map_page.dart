import 'dart:async';
import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:location/location.dart';
import 'package:client_user/app/providers.dart';
import 'package:client_user/app/home/models/shop.dart';
import 'package:latlng/latlng.dart' as latLng;

class MapPage extends StatelessWidget {
  const MapPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Center(
      child: MapsDemo(),
    ));
  }
}

void viewMap() {
  runApp(MaterialApp(
    home: Scaffold(
      body: ProviderScope(child: MapsDemo()),
    ),
  ));
}

class MapsDemo extends ConsumerStatefulWidget {
  @override
  MapsDemoState createState() => MapsDemoState();
}

// final customerProvider = FutureProvider<Shops?>(
//   (ref) => ref.read(databaseProvider)!.searchShop(_yourLocation),
// );

class MapsDemoState extends ConsumerState<MapsDemo> {
  late FutureProvider<Shops?> customerProvider;
  // final customerProvider = FutureProvider<Shops?>(
  //   (ref) => ref.read(databaseProvider)!.searchShop(_yourLocation),
  // );
  Completer<GoogleMapController> _controller = Completer();
  Location _locationService = Location();

  // 現在位置
  LocationData? _yourLocation;

  // 現在位置の監視状況
  StreamSubscription? _locationChangedListen;

  @override
  void initState() {
    super.initState();

    // 現在位置の取得
    _getLocation();

    customerProvider = FutureProvider<Shops?>(
      (ref) => ref.read(databaseProvider)!.searchShop(_yourLocation),
    );
    // 現在位置の変化を監視
    _locationChangedListen =
        _locationService.onLocationChanged.listen((LocationData result) async {
      setState(() {
        _yourLocation = result;
      });
    });
  }

  @override
  void dispose() {
    super.dispose();

    // 監視を終了
    _locationChangedListen?.cancel();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(children: [
        ElevatedButton(
            onPressed: () {
              ref.refresh(customerProvider);
            },
            child: const Text("Reload")),
        Expanded(child: _makeGoogleMap())
      ]),
    );
  }

  Widget _makeGoogleMap() {
    if (_yourLocation == null) {
      // 現在位置が取れるまではローディング中
      return Center(
        child: CircularProgressIndicator(),
      );
    } else {
      final _customerProvider = ref.watch(customerProvider);
      return _customerProvider.when(data: (data) {
        return GoogleMap(
          markers: data?.shops?.map((Shop shop) {
                return Marker(
                  markerId: MarkerId(shop.uid!),
                  position: LatLng(shop.latitude!, shop.longitude!),
                  infoWindow: InfoWindow(
                      title: shop.name,
                      snippet: shop.description,
                      onTap: () {
                        print("a");
                      }),
                );
              }).toSet() ??
              {},

          // 初期表示される位置情報を現在位置から設定
          initialCameraPosition: CameraPosition(
            target: LatLng(_yourLocation!.latitude!, _yourLocation!.longitude!),
            zoom: 18.0,
          ),
          onMapCreated: (GoogleMapController controller) {
            if (!_controller.isCompleted) {
              _controller.complete(controller);
            }
          },

          // 現在位置にアイコン（青い円形のやつ）を置く
          myLocationEnabled: true,
        );
      }, loading: () {
        return const CircularProgressIndicator();
      }, error: (error, stackTrace) {
        return const CircularProgressIndicator();
      });
    }
  }

  void _getLocation() async {
    _yourLocation = await _locationService.getLocation();
  }
}

// TODO: 明日の朝やる
// void displayShop() {
//   Navigator.of(context).push(MaterialPageRoute(
//       builder: (context) => ReceiptDialogPage(receipt: receipt, menus: menus)));
// }
