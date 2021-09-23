import 'dart:async';
import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:location/location.dart';

class MapPage extends StatelessWidget {
  const MapPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
        body: Center(
      child: ElevatedButton(child: const Text("地図を表示"), onPressed: viewMap),
    ));
  }
}

void viewMap() {
  runApp(MaterialApp(
    home: Scaffold(
      appBar: AppBar(title: const Text('Flutter Google Maps')),
      body: MapsDemo(),
    ),
  ));
}

class MapsDemo extends StatefulWidget {
  @override
  State createState() => MapsDemoState();
}

// class MapsDemoState extends State<MapsDemo> {
//   late GoogleMapController mapController;
//   Location _locationService = Location();
//   // 現在位置
//   LocationData? _yourLocation;
//   // 現在位置の監視状況
//   StreamSubscription? _locationChangedListen;
//   @override
//   void initState() {
//     super.initState();

//     // 現在位置の取得
//     _getLocation();

//     // 現在位置の変化を監視
//     _locationChangedListen =
//         _locationService.onLocationChanged.listen((LocationData result) async {
//       setState(() {
//         _yourLocation = result;
//       });
//     });
//   }

//   @override
//   void dispose() {
//     super.dispose();

//     // 監視を終了
//     _locationChangedListen?.cancel();
//   }

//   @override
//   Widget build(BuildContext context) {
//     // return Scaffold(
//     //   if (_yourLocation == null) {
//     //   // 現在位置が取れるまではローディング中
//     //   return Center(
//     //     child: CircularProgressIndicator(),
//     //   );
//     // } else {
//     //   // ここを追加
//     //   body: Container(
//     //     height: MediaQuery.of(context).size.height,
//     //     width: MediaQuery.of(context).size.width,
//     //     child: GoogleMap(
//     //         onMapCreated: _onMapCreated,
//     //         myLocationEnabled: true,
//     //         initialCameraPosition: CameraPosition(
//     //           //target: LatLng(35.6580339, 139.7016358),
//     //           target: LatLng(_yourLocation.latitude, _yourLocation.longitude),
//     //           zoom: 17.0,
//     //         )),
//     //   ),
//     // });
//     return Scaffold(
//       body: _makeGoogleMap(),
//     );
//   }

//   Widget _makeGoogleMap() {
//     if (_yourLocation == null) {
//       // 現在位置が取れるまではローディング中
//       return Center(
//         child: CircularProgressIndicator(),
//       );
//     } else {
//       // Google Map ウィジェットを返す
//       return GoogleMap(
//         // 初期表示される位置情報を現在位置から設定
//         initialCameraPosition: CameraPosition(
//           target: LatLng(_yourLocation.latitude, _yourLocation.longitude),
//           zoom: 18.0,
//         ),
//         onMapCreated: (GoogleMapController controller) {
//           _controller.complete(controller);
//         },

//         // 現在位置にアイコン（青い円形のやつ）を置く
//         myLocationEnabled: true,
//       );
//     }
//   }

//   void _getLocation() async {
//     _yourLocation = await _locationService.getLocation();
//   }

//   void _onMapCreated(GoogleMapController controller) {
//     setState(() {
//       mapController = controller;
//       mapController.animateCamera(CameraUpdate.newCameraPosition(
//         const CameraPosition(
//           target: LatLng(35.6580339, 139.7016358),
//           zoom: 17.0,
//         ),
//       ));
//     });
//   }
// }

class MapsDemoState extends State<MapsDemo> {
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
      body: _makeGoogleMap(),
    );
  }

  Widget _makeGoogleMap() {
    if (_yourLocation == null) {
      // 現在位置が取れるまではローディング中
      return Center(
        child: CircularProgressIndicator(),
      );
    } else {
      // Google Map ウィジェットを返す
      return GoogleMap(
        // 初期表示される位置情報を現在位置から設定
        initialCameraPosition: CameraPosition(
          target: LatLng(_yourLocation!.latitude!, _yourLocation!.longitude!),
          zoom: 18.0,
        ),
        onMapCreated: (GoogleMapController controller) {
          _controller.complete(controller);
        },

        // 現在位置にアイコン（青い円形のやつ）を置く
        myLocationEnabled: true,
      );
    }
  }

  void _getLocation() async {
    _yourLocation = await _locationService.getLocation();
  }
}
