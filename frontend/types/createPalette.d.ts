import { PaletteColor, PaletteColorOptions } from '@mui/material/styles/createPalette';

declare module '@mui/material/styles/createPalette' {
    interface Palette {
        dark?: PaletteColor;
        accent?: PaletteColor;
    }

    interface PaletteOptions {
        dark?: PaletteColorOptions;
        accent?: PaletteColorOptions;
    }
}
