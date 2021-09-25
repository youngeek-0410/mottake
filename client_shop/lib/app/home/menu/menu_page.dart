import 'package:client_shop/app/home/menu/menu_add_page.dart';
import 'package:client_shop/app/home/models/menu.dart';
import 'package:client_shop/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'menu_edit_page.dart';

final menuProvider = FutureProvider<MenuList>(
  (ref) => ref.read(databaseProvider)!.getMenus(),
);

class MenuPage extends ConsumerWidget {
  const MenuPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final _menuProvider = ref.watch(menuProvider);
    return Scaffold(
      body: Center(
        child: Column(
          children: [
            _menuProvider.when(data: (data) {
              if (data.menus != null) {
                return ListView(shrinkWrap: true, children: [
                  ...data.menus!
                      .map((e) => _listItem(
                              e.name!, e.price!, Icon(Icons.food_bank), () {
                            Navigator.of(context)
                                .push(MaterialPageRoute(
                                    builder: (context) => MenuEditPage(e)))
                                .then((_) => ref.refresh(menuProvider));
                          }))
                      .toList(),
                ]);
              } else {
                return const Text("No data");
              }
            }, loading: () {
              return const CircularProgressIndicator();
            }, error: (error, stackTrace) {
              return const Text("Error: could not fetch data");
            })
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.of(context)
              .push(MaterialPageRoute(builder: (_) => MenuAddPage()))
              .then((_) => ref.refresh(menuProvider));
        },
        child: const Icon(Icons.add),
      ),
    );
  }

  Widget _listItem(String title, int price, Icon icon, Function onTap) {
    return ListTile(
      title: Text(title),
      leading: icon,
      trailing: Text(price.toString() + " yen"),
      onTap: () {
        onTap();
      },
    );
  }
}
