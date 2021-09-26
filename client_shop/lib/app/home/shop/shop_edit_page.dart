import 'package:client_shop/app/home/common_widget.dart';
import 'package:client_shop/app/home/models/shop.dart';
import 'package:client_shop/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class ShopEditPage extends ConsumerStatefulWidget {
  final Shop shop;

  const ShopEditPage(this.shop, {Key? key}) : super(key: key);
  @override
  _ShopEditPageState createState() => _ShopEditPageState();
}

class _ShopEditPageState extends ConsumerState<ShopEditPage> {
  final _formKey = GlobalKey<FormState>();
  late String _name;
  late String _address;
  late String _description;
  late int _salesGoal;

  @override
  void initState() {
    super.initState();
    _name = widget.shop.name!;
    _address = widget.shop.address!;
    _description = widget.shop.description!;
    _salesGoal = widget.shop.salesGoal!;
  }

  @override
  Widget build(BuildContext contex) {
    return Scaffold(
        appBar: AppBar(
          title: const Text("Shop Editing Page"),
        ),
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
              _update();
            },
            child: const Text("Confirm")),
      ]),
    );
  }

  Future<void> _update() async {
    if (_validateAndSaveForm()) {
      try {
        final database = ref.read(databaseProvider)!;
        await database.updateShop(Shop(
            name: _name,
            address: _address,
            description: _description,
            salesGoal: _salesGoal));
        Navigator.of(context).pop();
      } catch (e) {
        showDialog(
            context: context,
            builder: (_) {
              return errorDialog("Failed to update menu.", () {
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
        decoration: const InputDecoration(labelText: 'ShopName'),
        keyboardAppearance: Brightness.light,
        initialValue: _name,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _name = value!,
      ),
      TextFormField(
        decoration: const InputDecoration(labelText: 'Address'),
        keyboardAppearance: Brightness.light,
        initialValue: _address,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _address = value!,
      ),
      TextFormField(
          decoration: const InputDecoration(labelText: 'Description'),
          keyboardAppearance: Brightness.light,
          initialValue: _description,
          minLines: 4,
          maxLines: null,
          keyboardType: TextInputType.multiline,
          onSaved: (value) => _description = value!),
      TextFormField(
        decoration: const InputDecoration(labelText: 'Sales Goal'),
        keyboardAppearance: Brightness.light,
        initialValue: _salesGoal.toString(),
        keyboardType: const TextInputType.numberWithOptions(
          signed: false,
          decimal: false,
        ),
        validator: (value) {
          final parsed = int.tryParse(value ?? '') ?? 0;
          return parsed > 0 ? null : 'Price must be greater than 0 yen';
        },
        onSaved: (value) => _salesGoal = int.tryParse(value ?? '') ?? 0,
      ),
    ];
  }
}
