import 'package:client_shop/app/home/common_widget.dart';
import 'package:client_shop/app/home/models/menu.dart';
import 'package:client_shop/app/providers.dart';
import 'package:client_shop/constants/strings.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class MenuAddPage extends ConsumerStatefulWidget {
  @override
  _MenuAddPageState createState() => _MenuAddPageState();
}

class _MenuAddPageState extends ConsumerState<MenuAddPage> {
  final _formKey = GlobalKey<FormState>();
  int? _id;
  String? _name;
  int? _price;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(title: const Text(Strings.menuCreationPage)),
        body: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            children: [_buildForm()],
          ),
        ));
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
            child: const Text(Strings.register))
      ]),
    );
  }

  Future<void> _submit() async {
    if (_validateAndSaveForm()) {
      try {
        final database = ref.read(databaseProvider)!;
        await database.createMenu(Menu(name: _name, price: _price));
        Navigator.of(context).pop();
      } catch (e) {
        showDialog(
            context: context,
            builder: (_) {
              return errorDialog("Failed to create menu.", () {
                Navigator.of(context).pop();
              });
            });
      }
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
        decoration: const InputDecoration(labelText: Strings.menuName),
        keyboardAppearance: Brightness.light,
        initialValue: _name,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _name = value,
      ),
      TextFormField(
        decoration: const InputDecoration(labelText: Strings.price),
        keyboardAppearance: Brightness.light,
        initialValue: "",
        keyboardType: const TextInputType.numberWithOptions(
          signed: false,
          decimal: false,
        ),
        validator: (value) {
          final parsed = int.tryParse(value ?? '') ?? 0;
          return parsed > 0 ? null : 'Price must be greater than 0 yen';
        },
        onSaved: (value) => _price = int.tryParse(value ?? '') ?? 0,
      ),
    ];
  }
}
