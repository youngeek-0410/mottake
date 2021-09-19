import 'package:client_shop/app/home/models/manu.dart';
import 'package:client_shop/app/providers.dart';
import 'package:confirm_dialog/confirm_dialog.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class MenuEditPage extends ConsumerStatefulWidget {
  final Menu menu;

  const MenuEditPage(this.menu, {Key? key}) : super(key: key);
  @override
  _MenuDetailPageState createState() => _MenuDetailPageState();
}

class _MenuDetailPageState extends ConsumerState<MenuEditPage> {
  final _formKey = GlobalKey<FormState>();
  int? _id;
  String? _name;
  int? _price;

  @override
  void initState() {
    super.initState();
    _name = widget.menu.name;
    _id = widget.menu.id;
    _price = widget.menu.price;
  }

  @override
  Widget build(BuildContext contex) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Menu Editing Page"),
      ),
      body: _buildForm(),
    );
  }

  Widget _buildForm() {
    return Form(
      key: _formKey,
      child: Column(crossAxisAlignment: CrossAxisAlignment.stretch, children: [
        ..._buildFormChildren(),
        ElevatedButton(
            onPressed: () {
              _submit();
            },
            child: const Text("Confirm")),
        ElevatedButton(
            onPressed: () async {
              if (await confirm(context)) {
                _delete();
              }
            },
            child: const Text("Delete"),
            style: ElevatedButton.styleFrom(primary: Colors.red)),
      ]),
    );
  }

  Future<void> _delete() async {
    final database = ref.read(databaseProvider)!;
    database
        .deleteMenu(_id!)
        .then((_) => Navigator.of(context).pop())
        .catchError((_) => Navigator.of(context).pop());
  }

  Future<void> _submit() async {
    if (_validateAndSaveForm()) {
      final database = ref.read(databaseProvider)!;
      database
          .patchMenu(Menu(id: _id, name: _name, price: _price))
          .then((_) => Navigator.of(context).pop())
          .catchError((_) => Navigator.of(context).pop());
    }
  }

  bool _validateAndSaveForm() {
    final form = _formKey.currentState!;
    if (form.validate()) {
      form.save();
      return true;
    }
    return false;
  }

  List<Widget> _buildFormChildren() {
    return [
      TextFormField(
        decoration: const InputDecoration(labelText: 'Name'),
        keyboardAppearance: Brightness.light,
        initialValue: _name,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _name = value,
      ),
      TextFormField(
        decoration: const InputDecoration(labelText: 'Rate per hour'),
        keyboardAppearance: Brightness.light,
        initialValue: _price.toString(),
        keyboardType: const TextInputType.numberWithOptions(
          signed: false,
          decimal: false,
        ),
        onSaved: (value) => _price = int.tryParse(value ?? '') ?? 0,
      ),
    ];
  }
}
