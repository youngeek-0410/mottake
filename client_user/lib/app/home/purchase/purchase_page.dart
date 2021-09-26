import 'package:flutter/material.dart';
import 'package:client_user/app/home/common_widget.dart';
import 'package:client_user/app/home/models/shop.dart';
import 'package:client_user/app/home/models/menu.dart';
import 'package:client_user/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:qr_flutter/qr_flutter.dart';

class PurchasePage extends StatelessWidget {
  const PurchasePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
        body: Center(
      child: Text("Purchase"),
    ));
  }
}

class DisplayShop extends ConsumerStatefulWidget {
  final Shop shop;
  DisplayShop({required this.shop});
  @override
  DisplayShopState createState() => DisplayShopState();
}

class DisplayShopState extends ConsumerState<DisplayShop> {
  late final FutureProvider<MenuList?> menuProvider;
  late final Shop shop;
  @override
  void initState() {
    super.initState();
    shop = widget.shop;
    menuProvider = FutureProvider<MenuList?>(
        (ref) => ref.read(databaseProvider)!.getMenus(shop.uid!));
  }

  // final List<Menu> menus;
  // late final FutureProvider<MenuList?> menuProvider;
  //= FutureProvider<MenuList?>(
  //   (ref) => ref.read(databaseProvider)!.getMenus(shop.uid!),
  // );
  Map<int, int> shoppingcourt = {};

  @override
  Widget build(BuildContext context) {
    //final List<Menu> menus;
    // final menuProvider = FutureProvider<MenuList?>(
    //   (ref) => ref.read(databaseProvider)!.getMenus(shop.uid!),
    // );
    final _menuProvider = ref.watch(menuProvider);
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: Column(
          children: [
            _menuProvider.when(data: (data) {
              return ListView(
                shrinkWrap: true,
                children: [
                  _listShopData(shop.name!, Icon(Icons.shop)),
                  ...?(data?.menus?.map((menu) {
                    return _listMenu(menu.name!, menu.price!, menu.id!);
                  }).toList()),
                  ElevatedButton(
                      onPressed: () {
                        _decidePurchase(context, ref);
                      },
                      child: const Text("購入"))
                ],
              );
            }, loading: () {
              return const CircularProgressIndicator();
            }, error: (error, stackTrace) {
              return const CircularProgressIndicator();
            })
          ],
        ),
      ),
    );
  }

  Widget _listMenu(String title, int price, int id) {
    bool _flag = false;
    if (shoppingcourt[id] == null) {
      shoppingcourt.addAll({id: 0});
    }
    // void _handleCheckbox(bool e) {
    //     _flag = e;
    // }
    return Card(
      elevation: 1.0,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: <Widget>[
            Expanded(
                child: ListTile(
              title: Text(title),
              trailing: Text("$price 円"),
            )),
            _decrementButton(id),
            Text(
              '${shoppingcourt[id]}',
              style: TextStyle(fontSize: 18.0),
            ),
            _incrementButton(id),
          ],
        ),
      ),
    );
    // return ListTile(
    //   title: Text(title),
    //   trailing: Text("$price yen"),
    // );
  }

  Widget _incrementButton(int id) {
    return FloatingActionButton(
      child: Icon(Icons.add, color: Colors.black87),
      backgroundColor: Colors.white,
      onPressed: () {
        setState(() {
          shoppingcourt[id] = shoppingcourt[id]! + 1;
        });
      },
    );
  }

  Widget _decrementButton(int id) {
    return FloatingActionButton(
        onPressed: () {
          setState(() {
            if (shoppingcourt[id]! > 0) {
              shoppingcourt[id] = shoppingcourt[id]! - 1;
            }
          });
        },
        child: Icon(Icons.remove, color: Colors.black87),
        backgroundColor: Colors.white);
  }

  Widget _listShopData(String title, Icon icon) {
    return ListTile(title: Text(title), leading: icon);
  }

  void _decidePurchase(BuildContext context, WidgetRef ref) async {
    // try {
    //   //final database = ref.read(databaseProvider)!;
    //   // await database.decidePurchase(receipt);
    //   Navigator.of(context).pop();
    // } catch (e) {
    //   await showDialog(
    //       context: context,
    //       builder: (_) => errorDialog("Failed to register receipt", () {
    //             Navigator.of(context).pop();
    //           }));
    //   Navigator.of(context).pop();
    // }
    final qrcode = ref.read(databaseProvider)!.generateQRcode(shoppingcourt);
    Navigator.of(context).push(MaterialPageRoute(builder: (_) {
      return QrCodeLayout(qrCode: qrcode);
    }));
  }
}

class QrCodeLayout extends StatelessWidget {
  const QrCodeLayout({Key? key, required this.qrCode}) : super(key: key);
  final String qrCode;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('QR Code Demo'),
      ),
      body: Center(
        child: QrImage(
          data: qrCode,
          size: 200,
        ),
      ),
    );
  }
}
