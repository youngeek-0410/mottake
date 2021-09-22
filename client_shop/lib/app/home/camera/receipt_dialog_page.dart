import 'package:client_shop/app/home/common_widget.dart';
import 'package:client_shop/app/home/models/manu.dart';
import 'package:client_shop/app/home/models/receipt.dart';
import 'package:client_shop/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class ReceiptDialogPage extends ConsumerWidget {
  final Receipt receipt;
  final List<Menu> menus;

  const ReceiptDialogPage({required this.receipt, required this.menus});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: Column(
          children: [
            ListView(
              shrinkWrap: true,
              children: [
                ...menus.map((menu) {
                  final number = receipt.purchases!
                      .where((purchase) => purchase.menuId! == menu.id!)
                      .first
                      .number;
                  return _listItem(menu.name!, menu.price!, number!);
                }).toList(),
                ElevatedButton(
                    onPressed: () {
                      _registerReceipt(context, ref);
                    },
                    child: const Text("Confirm"))
              ],
            )
          ],
        ),
      ),
    );
  }

  Widget _listItem(String title, int price, int number) {
    return ListTile(
      title: Text(title),
      trailing: Text("$price × $number yen"),
    );
  }

  void _registerReceipt(BuildContext context, WidgetRef ref) async {
    try {
      final database = ref.read(databaseProvider)!;
      await database.registerReceipt(receipt);
      Navigator.of(context).pop();
    } catch (e) {
      await showDialog(
          context: context,
          builder: (_) => errorDialog("Failed to register receipt", () {
                Navigator.of(context).pop();
              }));
      Navigator.of(context).pop();
    }
  }
}
