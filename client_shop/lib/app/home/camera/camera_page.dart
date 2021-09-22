import 'dart:convert';

import 'package:client_shop/app/home/camera/receipt_dialog_page.dart';
import 'package:client_shop/app/home/common_widget.dart';
import 'package:client_shop/app/home/models/menu.dart';
import 'package:client_shop/app/home/models/receipt.dart';
import 'package:client_shop/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_barcode_scanner/flutter_barcode_scanner.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class CameraPage extends ConsumerWidget {
  const CameraPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.camera),
        onPressed: () {
          FlutterBarcodeScanner.scanBarcode(
                  "#000000", "Cancel", false, ScanMode.DEFAULT)
              .then((qrcode) {
            if (qrcode != "-1") {
              _showDialog(qrcode, context, ref);
            }
          });
        },
      ),
    );
  }

  Future<void> _showDialog(
      String qrcode, BuildContext context, WidgetRef ref) async {
    try {
      final receipt = Receipt.fromJson(jsonDecode(qrcode));
      var menus = <Menu>[];
      for (var purchase in receipt.purchases!) {
        final menu =
            await ref.read(databaseProvider)!.getMenu(purchase.menuId!);
        menus.add(menu);
      }
      Navigator.of(context).push(MaterialPageRoute(
          builder: (context) =>
              ReceiptDialogPage(receipt: receipt, menus: menus)));
    } catch (e) {
      showDialog(
          context: context,
          builder: (_) {
            return errorDialog(
                "Failed to read data.", () => Navigator.of(context).pop());
          });
    }
  }
}
