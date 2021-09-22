import 'package:client_user/app/home/models/customer.dart';
import 'package:client_user/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:client_user/constants/urls.dart';

final customerProvider = FutureProvider<Customer>(
  (ref) => ref.read(databaseProvider)!.getCustomer(),
);

class CustomerPage extends ConsumerWidget {
  const CustomerPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final auth = ref.watch(authServiceProvider);
    final _customerProvider = ref.watch(customerProvider);
    return Scaffold(
        body: Center(
            child: Column(
      children: [
        _customerProvider.when(data: (data) {
          return ListView(
            shrinkWrap: true,
            children: [_listItem(data.name!, Icon(Icons.person))],
          );
        }, loading: () {
          return const CircularProgressIndicator();
        }, error: (error, stackTrace) {
          return const CircularProgressIndicator();
        }),
        ElevatedButton(
          child: const Text("Sign out"),
          onPressed: auth.signOut,
        ),
        ElevatedButton(
          child: const Text("Delete User"),
          onPressed: () {
            showDialog(
              context: context,
              builder: (_) {
                return AlertDialog(
                  title: Text("警告"),
                  content: Text("本当にユーザを削除しますか？この作業は取り消せません。"),
                  actions: <Widget>[
                    // ボタン領域
                    FlatButton(
                      child: Text("Cancel"),
                      onPressed: () => Navigator.pop(context),
                    ),
                    FlatButton(
                      child: Text("OK"),
                      onPressed: () async {
                        String url = URLs.baseURL + "/user";

                        http.Response resp = await http.delete(Uri.parse(url));
                        if (resp.statusCode != 200) {
                          setState(() {
                            int statusCode = resp.statusCode;
                            _content = "Failed to delete $statusCode";
                          });
                          return;
                        }
                        setState(() {
                          _content = resp.body;
                        });
                      },
                    ),
                  ],
                );
              },
            );
          },
        ),
      ],
    )));
  }

  Widget _listItem(String title, Icon icon) {
    return ListTile(title: Text(title), leading: icon);
  }
}
