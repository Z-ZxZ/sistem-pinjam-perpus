'use client';

import React from 'react';
import { motion } from 'framer-motion';
import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'accent' | 'outline' | 'ghost';
  isLoading?: boolean;
}

export const Button = ({ 
  className, 
  variant = 'primary', 
  isLoading, 
  children, 
  ...props 
}: ButtonProps) => {
  const variants = {
    primary: 'bg-[#4338CA] text-white hover:bg-[#3730A3]',
    accent: 'bg-[#14B8A6] text-white hover:bg-[#0D9488]',
    outline: 'border border-[#E2E8F0] text-[#1E293B] hover:bg-[#F8FAFC]',
    ghost: 'text-[#64748B] hover:bg-[#F8FAFC]',
  };

  return (
    <motion.button
      whileTap={{ scale: 0.98 }}
      className={twMerge(
        'px-6 py-2.5 rounded-lg font-medium transition-all duration-300 disabled:opacity-50 flex items-center justify-center gap-2',
        variants[variant],
        className
      )}
      disabled={isLoading || props.disabled}
      {...(props as any)}
    >
      {isLoading ? (
        <div className="w-5 h-5 border-2 border-current border-t-transparent rounded-full animate-spin" />
      ) : null}
      {children}
    </motion.button>
  );
};
