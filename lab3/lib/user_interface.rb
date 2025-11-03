require 'libui'
require_relative 'playfair_operations'

class UserInterface
  UI = LibUI

  def initialize
    UI.init

    window = UI.new_window('Lab 2', 400, 300, 1)
    UI.window_on_closing(window) { UI.quit; 0 }

    vbox = UI.new_vertical_box
    UI.window_set_child(window, vbox)

    toggle = UI.new_combobox
    UI.combobox_append(toggle, 'Encrypt')
    UI.combobox_append(toggle, 'Decrypt')
    UI.combobox_set_selected(toggle, 0)
    UI.box_append(vbox, toggle, 0)

    key_label = UI.new_label('Key:')
    UI.box_append(vbox, key_label, 0)
    key_input = UI.new_entry
    UI.box_append(vbox, key_input, 0)

    # text/cipher input field
    text_label = UI.new_label('Text/Cipher:')
    UI.box_append(vbox, text_label, 0)
    text_input = UI.new_multiline_entry
    UI.box_append(vbox, text_input, 1)

    output_label = UI.new_label('Output:')
    UI.box_append(vbox, output_label, 0)
    output_area = UI.new_multiline_entry
    UI.multiline_entry_set_read_only(output_area, 1)
    UI.box_append(vbox, output_area, 1)

    submit_button = UI.new_button('Submit')
    UI.button_on_clicked(submit_button) do
      mode = UI.combobox_selected(toggle) == 0 ? 'Encrypt' : 'Decrypt'
      key = UI.entry_text(key_input).to_s
      text = UI.multiline_entry_text(text_input).to_s

      # logic goes here
      playfair = PlayfairOperations.new

      case mode
      when 'Encrypt'
        output = playfair.encrypt(key, text)
      when 'Decrypt'
        output = playfair.decrypt(key, text)
      else
        output = 'Invalid operation mode'
      end

      UI.multiline_entry_set_text(output_area, output)
      0
    end
    UI.box_append(vbox, submit_button, 0)

    UI.window_set_margined(window, 1)
    UI.box_set_padded(vbox, 1)

    UI.control_show(window)
    UI.main
  end
end